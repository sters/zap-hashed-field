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

// Fnv128 generate zapcore.Field with Fnv128Hasher
func Fnv128(key string, value string) zapcore.Field {
	return String(key, Fnv128Hasher(value))
}
