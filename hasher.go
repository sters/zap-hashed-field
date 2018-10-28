package hashedfield

import (
	"crypto/sha1"
	"fmt"
)

// Hasher for use hashed field's value
type Hasher interface {
	Hash() string
}

// Sha1Hasher is a hasher using the sha1 method
type Sha1Hasher string

// Hash generate sha1 hashed string
func (v Sha1Hasher) Hash() string {
	h := sha1.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
