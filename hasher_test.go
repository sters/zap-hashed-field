package hashedfield_test

import (
	"testing"

	hashedfield "github.com/sters/zap-hashed-field"
)

func TestSha512Hasher_Hash(t *testing.T) {
	tests := []struct {
		name string
		v    hashedfield.Sha512Hasher
		want string
	}{
		{
			name: "HelloWorld",
			v:    hashedfield.Sha512Hasher("HelloWorld"),
			want: "8ae6ae71a75d3fb2e0225deeb004faf95d816a0a58093eb4cb5a3aa0f197050d7a4dc0a2d5c6fbae5fb5b0d536a0a9e6b686369fa57a027687c3630321547596",
		},
		{
			name: "blank",
			v:    hashedfield.Sha512Hasher(""),
			want: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Hash(); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFnv128Hasher_Hash(t *testing.T) {
	tests := []struct {
		name string
		v    hashedfield.Fnv128Hasher
		want string
	}{
		{
			name: "HelloWorld",
			v:    hashedfield.Fnv128Hasher("HelloWorld"),
			want: "c53c1eaca98cc15807e4f1a6db2d2be9",
		},
		{
			name: "blank",
			v:    hashedfield.Fnv128Hasher(""),
			want: "6c62272e07bb014262b821756295c58d",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Hash(); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
