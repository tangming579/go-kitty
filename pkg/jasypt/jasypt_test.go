package jasypt

import (
	"github.com/tangming579/go-kitty/internal/jasypt/iv"
	"github.com/tangming579/go-kitty/internal/jasypt/salt"
	"testing"
)

func TestJasypt(t *testing.T) {
	config := NewConfig(SetPassword("password"), SetSaltGenerator(salt.RandomSaltGenerator{}), SetIvGenerator(iv.NoIvGenerator{}))
	encryptor := New("PBEWithMD5AndDES", config)
	encrypt, err := encryptor.Encrypt(`plain text`)
	if err != nil {
		t.Error(err)
	}
	decrypt, err := encryptor.Decrypt(encrypt)
	if err != nil {
		t.Error(err)
	}
	if decrypt != `plain text` {
		t.Error(`decrypt failed`)
	}
}
