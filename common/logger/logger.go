// Package logger holds whole logic associated with logging information to console and files
package logger

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/DeRuina/timberjack"
)

func Setup() *slog.Logger {
	logRotator := timberjack.Logger{
		Filename:         "./logs/app.log",
		MaxSize:          100,
		MaxBackups:       7,
		MaxAge:           30,
		Compress:         true,
		RotationInterval: 24 * time.Hour,
	}

	return slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, &logRotator), nil))
}
