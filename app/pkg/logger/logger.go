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
type Struct struct {
	logger *zerolog.Logger
}

var _ Logger = (*Struct)(nil)

func New(level string) *Struct {
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

	return &Struct{
		logger: &logger,
	}
}
func NewLogger() Logger {
	return New("info")
}
func (l *Struct) Debug(msg string, args ...interface{}) {
	l.logger.Debug().Msgf(msg, args...)
}
func (l *Struct) Info(msg string, args ...interface{}) {
	l.logger.Info().Msgf(msg, args...)
}
func (l *Struct) Warn(msg string, args ...interface{}) {
	l.logger.Warn().Msgf(msg, args...)
}
func (l *Struct) Error(msg string, args ...interface{}) {
	l.logger.Error().Msgf(msg, args...)
}
func (l *Struct) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal().Msgf(msg, args...)
	os.Exit(1)
}
func (l *Struct) WithContext(_ context.Context) Logger {
	return &Struct{
		logger: l.logger,
	}
}
func (l *Struct) WithError(err error) Logger {
	newLogger := l.logger.With().Err(err).Logger()

	return &Struct{
		logger: &newLogger,
	}
}
func (l *Struct) WithField(key string, value interface{}) Logger {
	newLogger := l.logger.With().Interface(key, value).Logger()

	return &Struct{
		logger: &newLogger,
	}
}
func (l *Struct) WithFields(fields map[string]interface{}) Logger {
	newLogger := l.logger.With().Fields(fields).Logger()

	return &Struct{
		logger: &newLogger,
	}
}
func (l *Struct) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}
func (l *Struct) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}
func (l *Struct) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}
func (l *Struct) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}
func (l *Struct) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
	os.Exit(1)
}
