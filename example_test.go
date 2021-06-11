package hashedfield_test

import (
	hashedfield "github.com/sters/zap-hashed-field"
	"go.uber.org/zap"
)

func ExampleSha512() {
	logger := zap.NewExample()
	defer func() { _ = logger.Sync() }()

	key := "sensitive data"
	value := "Hello World!"

	logger.Info(
		"default output",
		zap.String(key, value),
	)

	logger.Info(
		"sha512 output",
		hashedfield.String(key, hashedfield.Sha512Hasher(value)),
	)

	logger.Info(
		"fnv128 output",
		hashedfield.String(key, hashedfield.Fnv128Hasher(value)),
	)

	// Unordered output:
	// {"level":"info","msg":"default output","sensitive data":"Hello World!"}
	// {"level":"info","msg":"sha512 output","sensitive data":"861844d6704e8573fec34d967e20bcfef3d424cf48be04e6dc08f2bd58c729743371015ead891cc3cf1c9d34b49264b510751b1ff9e537937bc46b5d6ff4ecc8"}
	// {"level":"info","msg":"fnv128 output","sensitive data":"3e09e68d3967a18bcfaa7ac886326144"}
}
