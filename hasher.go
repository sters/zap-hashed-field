package hashedfield

import (
	"crypto/sha512"
	"fmt"
	"hash/fnv"
)

// Hasher for use hashed field's value
type Hasher interface {
	Hash() string
}

// Sha512Hasher is a hasher using the sha1 method
type Sha512Hasher string

// Hash generate Sha512 hashed string
func (v Sha512Hasher) Hash() string {
	h := sha512.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Fnv128Hasher is a hasher using the sha1 method
type Fnv128Hasher string

// Hash generate Sha512 hashed string
func (v Fnv128Hasher) Hash() string {
	h := fnv.New128()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
