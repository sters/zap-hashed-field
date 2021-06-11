package hashedfield

import (
	"crypto/sha512"
	"fmt"
)

// Hasher for use hashed field's value
type Hasher interface {
	Hash() string
}

// Sha1Hasher is a hasher using the sha1 method
type Sha512Hasher string

// Hash generate Sha512 hashed string
func (v Sha512Hasher) Hash() string {
	h := sha512.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
