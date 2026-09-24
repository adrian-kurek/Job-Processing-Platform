package main

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/DeRuina/timberjack"
)

func main() {
	logRotator := timberjack.Logger{
		Filename:         "./logs/app.log",
		MaxSize:          100,
		MaxBackups:       7,
		MaxAge:           30,
		Compress:         true,
		RotationInterval: 24 * time.Hour,
	}

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, &logRotator), nil))
	logger.Info("Application started")
}
