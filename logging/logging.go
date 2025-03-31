package logging

import (
	"log"
	"log/slog"
	"os"
)

type ILogger interface {
	Info(msg string, params ...any)
	Error(msg string, params ...any)
	Fatal(msg string)
}

type ConsoleLogger struct {
	InfoLevel  *log.Logger
	ErrorLevel *log.Logger
}

func (l *ConsoleLogger) Info(msg string, params ...any) {
	slog.Info(msg, params...)
}

func (l *ConsoleLogger) Error(msg string, params ...any) {
	slog.Error(msg, params...)
}

func (l *ConsoleLogger) Fatal(msg string) {
	slog.Error(msg)

	os.Exit(1)
}
