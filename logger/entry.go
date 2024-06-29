package logger

import (
	"context"
	"fmt"

	composition "github.com/ginger-core/compound/context"
	"go.uber.org/zap/zapcore"
)

type Entry interface {
	Logger
}

type entry struct {
	*logger

	source   string
	id       string
	uid      string
	path     string
	severity string
}

func NewEntry(l Logger) Entry {
	e := &entry{
		logger: l.(*logger),
	}
	return e
}

func (l *entry) Clone() Logger {
	return l.clone()
}

func (l *entry) WithContext(ctx context.Context) Logger {
	data := composition.GetContextData(ctx)
	if len(data) == 0 {
		return l
	}
	f := make(Field)
	for k, v := range data {
		f[string(k)] = v
	}
	return l.With(f)
}

func (l *entry) With(field Field) Logger {
	l = l.clone()
	fields := l.getFields(field)
	for i, _l := range l.loggers {
		l.loggers[i] = _l.With(fields...)
	}
	return l
}

func (l *entry) WithSource(source string) Logger {
	l = l.clone()
	l.source = source
	return l
}

func (l *entry) SetSource(source string) {
	l.source = source
}

func (l *entry) WithId(id string) Logger {
	l = l.clone()
	l.id = id
	return l
}

func (l *entry) WithUid(uid string) Logger {
	l = l.clone()
	l.uid = uid
	return l
}

func (l *entry) WithTrace(path string) Logger {
	l = l.clone()
	if l.path != "" {
		l.path += ">"
	}
	l.path += path
	return l
}

func (l *entry) logf(level Level, msg string) {
	baseFields := []zapcore.Field{
		{
			Key:    "source",
			Type:   zapcore.StringType,
			String: l.source,
		},
		{
			Key:    "severity",
			Type:   zapcore.StringType,
			String: l.severity,
		},
		{
			Key:    "id",
			Type:   zapcore.StringType,
			String: l.id,
		},
		{
			Key:    "uid",
			Type:   zapcore.StringType,
			String: l.uid,
		},
		{
			Key:    "trace",
			Type:   zapcore.StringType,
			String: l.path,
		},
	}
	for _, logger := range l.loggers {
		if ce := logger.Check(GetZapLevel(level), msg); ce != nil {
			ce.Write(baseFields...)
		}
	}
}

func (l *entry) Debugf(format string, args ...any) {
	l.withSeverity(SeverityLow).
		logf(DebugLevel, fmt.Sprintf(format, args...))
}

func (l *entry) Infof(format string, args ...any) {
	l.withSeverity(SeverityLow).
		logf(InfoLevel, fmt.Sprintf(format, args...))
}

func (l *entry) Warnf(format string, args ...any) {
	l.withSeverity(SeverityMedium).
		logf(WarnLevel, fmt.Sprintf(format, args...))
}

func (l *entry) Errorf(format string, args ...any) {
	l.withSeverity(SeverityHigh).
		logf(ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *entry) Criticalf(format string, args ...any) {
	l.withSeverity(SeverityCritical).
		logf(ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *entry) Fatalf(format string, args ...any) {
	l.withSeverity(SeverityCritical).
		logf(FatalLevel, fmt.Sprintf(format, args...))
}

func (l *entry) withSeverity(severity Severity) *entry {
	l = l.clone()
	l.severity = string(severity)
	return l
}

func (l *entry) clone() *entry {
	cloned := &entry{
		logger:   l.logger.clone(),
		source:   l.source,
		id:       l.id,
		uid:      l.uid,
		path:     l.path,
		severity: l.severity,
	}
	return cloned
}
