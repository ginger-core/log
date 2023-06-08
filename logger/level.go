package logger

import "go.uber.org/zap/zapcore"

// Level is the minimum enabled logging level. Note that this is a dynamic
// level, so calling Config.Level.SetLevel will atomically change the log
// level of all loggers descended from this config.
type Level string

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = "debug"
	// InfoLevel is the default logging priority.
	InfoLevel Level = "info"
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel Level = "warn"
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel Level = "error"
	// PanicLevel logs a message, then panics.
	PanicLevel Level = "panic"
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = "fatal"
)

func GetZapLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
