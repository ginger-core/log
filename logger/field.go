package logger

import "go.uber.org/zap/zapcore"

type Field map[string]any

func (f Field) clone() Field {
	cloned := make(Field)
	for k, v := range f {
		cloned[k] = v
	}
	return cloned
}

func (l *entry) getFields(f Field) []zapcore.Field {
	fields := make([]zapcore.Field, 0)
	for k, f := range f {
		switch f.(type) {
		case string:
			fields = append(fields, zapcore.Field{String: f.(string), Type: zapcore.StringType, Key: k})
		case int:
			fields = append(fields, zapcore.Field{Integer: int64(f.(int)), Type: zapcore.Int64Type, Key: k})
		case error:
			fields = append(fields, zapcore.Field{Interface: f.(error), Type: zapcore.ErrorType, Key: k})
		case bool:
			fields = append(fields, zapcore.Field{Interface: f.(bool), Type: zapcore.BoolType, Key: k})
		default:
			fields = append(fields, zapcore.Field{Interface: f, Type: zapcore.ReflectType, Key: k})
		}
	}
	return fields
}
