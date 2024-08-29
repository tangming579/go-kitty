package jasypt

import (
	"github.com/tangming579/go-kitty/internal/jasypt/encryption"
	"github.com/tangming579/go-kitty/internal/jasypt/iv"
	"github.com/tangming579/go-kitty/internal/jasypt/salt"
	"regexp"
)

const (
	algorithm = "PBEWithMD5AndDES"
	encReg    = `ENC\(([^)]*)\)`
)

type basicTextEncryptor struct {
	encryption.Encryptor
}

func CreateBasicTextEncryptor(password string) basicTextEncryptor {
	config := NewConfig(SetPassword(password), SetSaltGenerator(salt.RandomSaltGenerator{}), SetIvGenerator(iv.NoIvGenerator{}))
	encryptor := New(algorithm, config)
	enc := basicTextEncryptor{encryptor}
	return enc
}

func (enc *basicTextEncryptor) Encrypt(message string) (string, error) {
	encrypt, err := enc.Encryptor.Encrypt(message)
	if err != nil {
		return message, err
	}
	return encrypt, nil
}

func (enc *basicTextEncryptor) Decrypt(message string) (string, error) {
	decrypt, err := enc.Encryptor.Decrypt(message)
	if err != nil {
		return message, err
	}
	return decrypt, nil
}

func (enc *basicTextEncryptor) GetString(original string) string {
	re := regexp.MustCompile(encReg)
	matches := re.FindStringSubmatch(original)
	if len(matches) == 0 {
		return original
	}
	message := matches[1]
	decrypt, err := enc.Decrypt(message)
	if err != nil {
		return message
	}
	return decrypt
}
