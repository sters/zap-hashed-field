package hashedfield_test

import (
	"reflect"
	"testing"

	hashedfield "github.com/sters/zap-hashed-field"
	"go.uber.org/zap/zapcore"
)

// nonhasher is dammy Hasher
type nonhasher string

func (n nonhasher) Hash() string {
	return string(n)
}

func TestString(t *testing.T) {
	type args struct {
		key   string
		value hashedfield.Hasher
	}

	// wants is simple expect zapcore.Field
	type wants struct {
		Key    string
		Type   zapcore.FieldType
		String string
	}

	text := "dammy"

	arg := args{
		key:   "dammy",
		value: nonhasher(text),
	}

	want := wants{
		Key:    arg.key,
		Type:   zapcore.StringType,
		String: text,
	}

	got := hashedfield.String(arg.key, arg.value)

	if !reflect.DeepEqual(got.Key, want.Key) {
		t.Errorf("String().Key = %v, want %v", got.Key, want.Key)
	}

	if !reflect.DeepEqual(got.Type, want.Type) {
		t.Errorf("String().Type = %v, want %v", got.Type, want.Type)
	}

	if !reflect.DeepEqual(got.String, want.String) {
		t.Errorf("String().String = %v, want %v", got.String, want.String)
	}
}

func TestSha1(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	// wants is simple expect zapcore.Field
	type wants struct {
		Key    string
		Type   zapcore.FieldType
		String string
	}

	text := "dammy"
	sha1text := "2b469fb45502c6d28aa4abef62d721961e27ef76"

	arg := args{
		key:   "dammy",
		value: text,
	}

	want := wants{
		Key:    arg.key,
		Type:   zapcore.StringType,
		String: sha1text,
	}

	got := hashedfield.Sha1(arg.key, arg.value)

	if !reflect.DeepEqual(got.Key, want.Key) {
		t.Errorf("Sha1().Key = %v, want %v", got.Key, want.Key)
	}

	if !reflect.DeepEqual(got.Type, want.Type) {
		t.Errorf("Sha1().Type = %v, want %v", got.Type, want.Type)
	}

	if !reflect.DeepEqual(got.String, want.String) {
		t.Errorf("Sha1().String = %v, want %v", got.String, want.String)
	}
}
