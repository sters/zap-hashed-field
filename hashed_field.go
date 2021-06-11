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

// Sha512 generate zapcore.Field with Sha512Hasher
func Sha512(key string, value string) zapcore.Field {
	return String(key, Sha512Hasher(value))
}
