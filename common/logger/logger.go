// Package logger holds whole logic associated with logging information to console and files
package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/DeRuina/timberjack"
)

type logSource string

const (
	APILogSource     logSource = "api"
	WorkerLogSource  logSource = "worker"
	maxFileSize                = 100
	maxBackups                 = 7
	maxAge                     = 30
	rotationInterval           = 24
)

func Setup(source logSource) *slog.Logger {
	logRotator := timberjack.Logger{
		Filename:         fmt.Sprintf("./logs/%s.log", source),
		MaxSize:          maxFileSize,
		MaxBackups:       maxBackups,
		MaxAge:           maxAge,
		RotationInterval: rotationInterval * time.Hour,
	}

	return slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, &logRotator), nil)).With("source", source)
}
