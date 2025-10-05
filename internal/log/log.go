package log

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type loggerContextKey struct{}

func WithLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, logger)
}

func FromContext(ctx context.Context) *zerolog.Logger {
	logger, ok := ctx.Value(loggerContextKey{}).(*zerolog.Logger)
	if !ok {
		nop := zerolog.Nop()
		return &nop
	}

	return logger
}

func LogLevel() zerolog.Level {
	levelEnvVar, _ := os.LookupEnv("LOG_LEVEL")

	switch strings.ToLower(levelEnvVar) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel

	default:
		return zerolog.WarnLevel
	}
}

func NewLogger(appName, appVersion string) (*zerolog.Logger, error) {
	// Use console writer for human-readable output
	output := zerolog.ConsoleWriter{Out: os.Stderr}

	logger := zerolog.New(output).
		Level(LogLevel()).
		With().
		Timestamp().
		Str("app_name", appName).
		Str("version", appVersion).
		Logger()

	return &logger, nil
}
