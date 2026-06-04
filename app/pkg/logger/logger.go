package logger

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	WithContext(ctx context.Context) Logger
	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
type AppLogger struct {
	logger *zerolog.Logger
}

var _ Logger = (*AppLogger)(nil)

func New(level string) *AppLogger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &AppLogger{
		logger: &logger,
	}
}
func NewLogger() Logger {
	return New("info")
}
func (l *AppLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debug().Msgf(msg, args...)
}
func (l *AppLogger) Info(msg string, args ...interface{}) {
	l.logger.Info().Msgf(msg, args...)
}
func (l *AppLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warn().Msgf(msg, args...)
}
func (l *AppLogger) Error(msg string, args ...interface{}) {
	l.logger.Error().Msgf(msg, args...)
}
func (l *AppLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal().Msgf(msg, args...)
	os.Exit(1)
}
func (l *AppLogger) WithContext(_ context.Context) Logger {
	return &AppLogger{
		logger: l.logger,
	}
}
func (l *AppLogger) WithError(err error) Logger {
	newLogger := l.logger.With().Err(err).Logger()

	return &AppLogger{
		logger: &newLogger,
	}
}
func (l *AppLogger) WithField(key string, value interface{}) Logger {
	newLogger := l.logger.With().Interface(key, value).Logger()

	return &AppLogger{
		logger: &newLogger,
	}
}
func (l *AppLogger) WithFields(fields map[string]interface{}) Logger {
	newLogger := l.logger.With().Fields(fields).Logger()

	return &AppLogger{
		logger: &newLogger,
	}
}
func (l *AppLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}
func (l *AppLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}
func (l *AppLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}
func (l *AppLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}
func (l *AppLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
	os.Exit(1)
}
