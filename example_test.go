package hashedfield_test

import (
	hashedfield "github.com/sters/zap-hashed-field"
	"go.uber.org/zap"
)

func ExampleSha1() {
	logger := zap.NewExample()
	defer logger.Sync()

	key := "sensitive data"
	value := "Hello World!"

	logger.Info(
		"default output",
		zap.String(key, value),
	)

	logger.Info(
		"sha1 output",
		hashedfield.Sha1(key, value),
	)

	// Unordered output:
	// {"level":"info","msg":"default output","sensitive data":"Hello World!"}
	// {"level":"info","msg":"sha1 output","sensitive data":"2ef7bde608ce5404e97d5f042f95f89f1c232871"}
}
