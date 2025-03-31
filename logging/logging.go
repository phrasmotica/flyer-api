package logging

import (
	"log"
)

type ILogger interface {
	Info(params ...any)
	Infof(format string, params ...any)
	Error(params ...any)
	Fatal(params ...any)
}

type ConsoleLogger struct {
	InfoLevel  *log.Logger
	ErrorLevel *log.Logger
}

func (l *ConsoleLogger) Info(params ...any) {
	l.InfoLevel.Println(params...)
}

func (l *ConsoleLogger) Infof(format string, params ...any) {
	l.InfoLevel.Printf(format, params...)
}

func (l *ConsoleLogger) Error(params ...any) {
	l.ErrorLevel.Println(params...)
}

func (l *ConsoleLogger) Fatal(params ...any) {
	l.ErrorLevel.Fatal(params...)
}
