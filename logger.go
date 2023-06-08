package log

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/log/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	logger.Logger
}

type Handler interface {
	Logger
	Start()
}

type _logger struct {
	logger.Logger
}

func NewLogger(registry registry.Registry) Handler {
	config := new(logger.Config)

	if err := registry.Unmarshal(config); err != nil {
		panic(err)
	}

	encoder := getEncoder(config.Encoding, zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   logger.ShortCallerEncoder,
	})

	l := logger.New()
	l.WithEntry(logger.NewEntry(l))

	writersConf := registry.ValueOf("writers")

	writers := make([]logger.Writer, 0)
	for w, _ := range config.Writers {
		wr := writer.New(w, writersConf.ValueOf(w))
		if wr == nil {
			continue
		}
		writers = append(writers, wr)
	}
	l.WithWriters(writers...)

	zapLevel := logger.GetZapLevel(config.Level)
	loggers := make([]*zap.Logger, 0)
	for _, w := range writers {
		level := zapLevel
		if lvl := w.Level(); lvl != "" {
			level = logger.GetZapLevel(lvl)
		}
		loggers = append(loggers,
			zap.New(
				zapcore.NewCore(
					encoder,
					zapcore.AddSync(w),
					zap.NewAtomicLevelAt(level),
				),
			),
		)
	}
	l.WithLoggers(loggers...)

	return &_logger{
		Logger: l,
	}
}

func getEncoder(encoding string, config zapcore.EncoderConfig) zapcore.Encoder {
	switch encoding {
	case "console":
		return zapcore.NewConsoleEncoder(config)
	case "json":
		return zapcore.NewJSONEncoder(config)
	}
	return nil
}
