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
	slogger *slog.Logger
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
	slog.SetDefault(defaultLogger.slogger)
	defaultLogger.Debug("log level set", "LOG_LEVEL", lvl.String())
}

func (l *Logger) Debug(msg string, args ...any) {
	l.slogger.Debug(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.slogger.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.slogger.Log(context.Background(), levelFatal, msg, args...)
	os.Exit(1)
}

func (l *Logger) Info(msg string, args ...any) {
	l.slogger.Info(msg, args...)
}

func (l *Logger) With(fields Fields) *Logger {
	args := []any{}
	for k, v := range fields {
		args = append(args, k, v)
	}
	return &Logger{l.slogger.With(args...)}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{l.slogger.With("err", err)}
}

func (l *Logger) WithGroup(name string) *Logger {
	return &Logger{l.slogger.WithGroup(name)}
}
