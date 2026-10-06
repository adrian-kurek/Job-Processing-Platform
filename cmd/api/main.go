package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/adrian-kurek/Job-Processing-Platform/common/logger"
	"github.com/adrian-kurek/Job-Processing-Platform/config"
	"github.com/adrian-kurek/Job-Processing-Platform/internal/job"
	"github.com/adrian-kurek/Job-Processing-Platform/internal/server"
	"github.com/joho/godotenv"
)

func connectToDB() (*config.DB, error) {
	dbConnectionLink := os.Getenv("DB_LINK")

	db, err := config.NewDB(dbConnectionLink)
	if err != nil {
		return &config.DB{}, err
	}
	return db, nil
}

func bootstrapDependencies(logger *slog.Logger, db *config.DB, port string) *server.HTTP {
	jobRepository := job.NewRepository(db.Connection, logger)
	jobService := job.NewService(jobRepository, logger)
	jobHandler := job.NewHandler(jobService, logger)
	dependencies := server.NewDependencyConfig(port, *jobHandler)
	return server.NewHTTP(dependencies)
}

func main() {
	logger := logger.Setup(logger.APILogSource)
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load environment variables", "err", err.Error())
		panic(err)
	}

	db, err := connectToDB()
	if err != nil {
		logger.Error("Failed to connecto to databse service", "err", err.Error())
		panic(err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			logger.Error("failed to close db connection", "err", err.Error())
		}
	}()

	apiCtx, apiCtxCancel := context.WithCancel(context.Background())
	port := os.Getenv("PORT")

	httpServer := bootstrapDependencies(logger, db, port)
	go func() {
		logger.Info("Applicattion started", "port", port)
		if err = httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Failed to start server", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer apiCtxCancel()

	if err = httpServer.Shutdown(apiCtx); err != nil {
		logger.Error("Server forced to shutdown", "err", err)
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}

	logger.Info("server exited")
}
