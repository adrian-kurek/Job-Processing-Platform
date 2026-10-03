package main

import (
	"os"

	"github.com/adrian-kurek/Job-Processing-Platform/common/logger"
	"github.com/adrian-kurek/Job-Processing-Platform/config"
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

func main() {
	logger := logger.Setup(logger.APILogSource)
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load environment variables", "err:", err.Error())
		panic(err)
	}

	db, err := connectToDB()
	if err != nil {
		logger.Error("Failed to connecto to databse service", "err:", err.Error())
		panic(err)
	}
	defer db.Close()

	logger.Info("Application started")
}
