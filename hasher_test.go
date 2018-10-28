package hashedfield_test

import (
	"testing"

	hashedfield "github.com/sters/zap-hashed-field"
)

func TestSha1Hasher_Hash(t *testing.T) {
	tests := []struct {
		name string
		v    hashedfield.Sha1Hasher
		want string
	}{
		{
			name: "HelloWorld",
			v:    hashedfield.Sha1Hasher("HelloWorld"),
			want: "db8ac1c259eb89d4a131b253bacfca5f319d54f2",
		},
		{
			name: "blank",
			v:    hashedfield.Sha1Hasher(""),
			want: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Hash(); got != tt.want {
				t.Errorf("Sha1Hasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
