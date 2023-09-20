package logger

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

const (
	levelFatal = slog.Level(12)
)

var levelNames = map[slog.Leveler]string{
	levelFatal: "FATAL",
}

var defaultLogger *Logger

type Fields map[string]any

type Logger struct {
	logger *slog.Logger
}

func Initialise(debug bool) {
	lvl := slog.LevelInfo
	if debug {
		lvl = slog.LevelDebug
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	})
	defaultLogger = &Logger{slog.New(handler)}
	slog.SetDefault(defaultLogger.logger)
	defaultLogger.Debug("log level set", "LOG_LEVEL", lvl.String())
}

// Internal Logger Funcs

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.logger.Log(context.Background(), levelFatal, msg, args...)
	os.Exit(1)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) With(fields Fields) *Logger {
	args := []any{}
	for k, v := range fields {
		args = append(args, k, v)
	}
	return &Logger{l.logger.With(args...)}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{l.logger.With("err", err)}
}

func (l *Logger) WithGroup(name string) *Logger {
	return &Logger{l.logger.WithGroup(name)}
}

// External Logger Funcs

func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	defaultLogger.Fatal(msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

func With(fields Fields) *Logger {
	return defaultLogger.With(fields)
}

func WithError(err error) *Logger {
	return defaultLogger.WithError(err)
}

func WithGroup(name string) *Logger {
	return defaultLogger.WithGroup(name)
}
