package logger

import (
	"log"
	"os"
)

// Logger interface is used by the SDK to log messages.
// Should be used to provide custom logging writers for the SDK to use.
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Error(...interface{})
	SetLevel(level Level)
}

// NewDefaultLogger returns a Logger which will write log messages to stdout,
// and use same formatting runes as the stdlib log.Logger
func NewDefaultLogger() Logger {
	return &defaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  LevelInfo,
	}
}

// A defaultLogger provides a minimalistic logger satisfying the Logger interface.
type defaultLogger struct {
	level  Level
	logger *log.Logger
}

func (l defaultLogger) Debug(args ...interface{}) {
	if l.shouldLog(LevelDebug) {
		l.log("[level:DEBUG]", args)
	}
}

func (l defaultLogger) Info(args ...interface{}) {
	if l.shouldLog(LevelInfo) {
		l.log("[level:INFO]", args)
	}
}

func (l defaultLogger) Error(args ...interface{}) {
	if l.shouldLog(LevelError) {
		l.log("[level:ERROR]", args)
	}
}

// Log logs the parameters to the stdlib logger. See log.Println.
func (l defaultLogger) log(args ...interface{}) {
	l.logger.Println(args...)
}
