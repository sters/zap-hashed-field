package hashedfield

import (
	"go.uber.org/zap/zapcore"
)

// String generate zapcore.Field with your Hasher
func String(key string, value Hasher) zapcore.Field {
	return zapcore.Field{
		Key:    key,
		Type:   zapcore.StringType,
		String: value.Hash(),
	}
}

// Sha1 generate zapcore.Field with Sha1Hasher
func Sha1(key string, value string) zapcore.Field {
	return String(key, Sha1Hasher(value))
}
