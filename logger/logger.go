package logger

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	Clone() Logger

	WithContext(ctx context.Context) Logger

	With(field Field) Logger
	WithSource(source string) Logger
	SetSource(source string)
	// WithId unique id of log
	//
	// Deprecated: use WithTrace instead
	WithId(id string) Logger
	WithUid(uid string) Logger
	WithTrace(path string) Logger

	WithWriters(writers ...Writer) Logger
	GetWriters() []Writer

	WithLoggers(loggers ...*zap.Logger) Logger
	GetLoggers() []*zap.Logger

	WithEntry(entry Entry) Logger
	GetEntry() Entry

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Criticalf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type logger struct {
	*entry
	writers []Writer
	loggers []*zap.Logger
}

func New() Logger {
	return &logger{
		writers: make([]Writer, 0),
		loggers: make([]*zap.Logger, 0),
	}
}

func (l *logger) Clone() Logger {
	cloned := &logger{
		entry:   l.entry,
		writers: l.writers,
		loggers: make([]*zap.Logger, len(l.loggers)),
	}
	return cloned
}

func (l *logger) clone(fields ...zap.Field) *logger {
	cloned := &logger{
		entry:   l.entry,
		writers: l.writers,
		loggers: make([]*zap.Logger, len(l.loggers)),
	}
	for i, _l := range l.loggers {
		cloned.loggers[i] = _l.With(fields...)
	}
	return cloned
}

func (l *logger) WithWriters(writers ...Writer) Logger {
	l.writers = append(l.writers, writers...)
	return l
}

func (l *logger) GetWriters() []Writer {
	return l.writers
}

func (l *logger) WithLoggers(loggers ...*zap.Logger) Logger {
	l.loggers = append(l.loggers, loggers...)
	return l
}

func (l *logger) GetLoggers() []*zap.Logger {
	return l.loggers
}

func (l *logger) WithEntry(e Entry) Logger {
	l.entry = e.(*entry)
	return l
}

func (l *logger) GetEntry() Entry {
	return l.entry
}
