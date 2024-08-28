package encryption

import (
	"github.com/tangming579/go-kitty/internal/jasypt/iv"
	"github.com/tangming579/go-kitty/internal/jasypt/salt"
)

type Encryptor interface {
	Encrypt(message string) (string, error)
	Decrypt(message string) (string, error)
}

type EncryptorConfig struct {
	Password      string
	SaltGenerator salt.Generator
	IvGenerator   iv.Generator
}
