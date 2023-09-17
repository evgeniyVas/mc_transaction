package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	stdOutLogger *slog.Logger
}

var onceLog *Logger

func NewLogger() {
	onceLog = newLogger()
}

func newLogger() *Logger {
	var level slog.Level
	if os.Getenv("APP_MODE") != "prod" {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{Level: level})
	l := slog.New(handler)

	return &Logger{stdOutLogger: l}
}

func Info(msg string, args ...any) {
	onceLog.stdOutLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	onceLog.stdOutLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	onceLog.stdOutLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	onceLog.stdOutLogger.Debug(msg, args...)
}
