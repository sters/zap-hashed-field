package hashedfield_test

import (
	"reflect"
	"testing"

	hashedfield "github.com/sters/zap-hashed-field"
	"go.uber.org/zap/zapcore"
)

// nonhasher is dummy Hasher
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

func TestSha512(t *testing.T) {
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
	hashedText := "5a2a2b483083d88baa51cfc1e3f62efc2ce9d27baa9e773cbea9d154f5fca79954ba8afdeea88c06ea33f642743caa491b4a25446a64477cd71ae30c8db74caa"

	arg := args{
		key:   "dammy",
		value: text,
	}

	want := wants{
		Key:    arg.key,
		Type:   zapcore.StringType,
		String: hashedText,
	}

	got := hashedfield.Sha512(arg.key, arg.value)

	if !reflect.DeepEqual(got.Key, want.Key) {
		t.Errorf("Key = %v, want %v", got.Key, want.Key)
	}

	if !reflect.DeepEqual(got.Type, want.Type) {
		t.Errorf("Type = %v, want %v", got.Type, want.Type)
	}

	if !reflect.DeepEqual(got.String, want.String) {
		t.Errorf("String = %v, want %v", got.String, want.String)
	}
}

func TestFnv128(t *testing.T) {
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
	hashedText := "f16a5b600583d94f7080388335d98197"

	arg := args{
		key:   "dammy",
		value: text,
	}

	want := wants{
		Key:    arg.key,
		Type:   zapcore.StringType,
		String: hashedText,
	}

	got := hashedfield.Fnv128(arg.key, arg.value)

	if !reflect.DeepEqual(got.Key, want.Key) {
		t.Errorf("Key = %v, want %v", got.Key, want.Key)
	}

	if !reflect.DeepEqual(got.Type, want.Type) {
		t.Errorf("Type = %v, want %v", got.Type, want.Type)
	}

	if !reflect.DeepEqual(got.String, want.String) {
		t.Errorf("String = %v, want %v", got.String, want.String)
	}
}
