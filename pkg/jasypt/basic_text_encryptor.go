package jasypt

import (
	"github.com/tangming579/go-kitty/internal/jasypt/iv"
	"github.com/tangming579/go-kitty/internal/jasypt/salt"
	"regexp"
)

func Encrypt(password, message string) (string, error) {
	config := NewConfig(SetPassword(password), SetSaltGenerator(salt.RandomSaltGenerator{}), SetIvGenerator(iv.NoIvGenerator{}))
	encryptor := New("PBEWithMD5AndDES", config)
	encrypt, err := encryptor.Encrypt(message)
	if err != nil {
		return message, err
	}
	return encrypt, nil
}

func Decrypt(password, message string) (string, error) {
	config := NewConfig(SetPassword(password), SetSaltGenerator(salt.RandomSaltGenerator{}), SetIvGenerator(iv.NoIvGenerator{}))
	encryptor := New("PBEWithMD5AndDES", config)
	decrypt, err := encryptor.Decrypt(message)
	if err != nil {
		return message, err
	}
	return decrypt, nil
}

func GetString(password, original string) string {
	re := regexp.MustCompile(`ENC\(([^)]*)\)`)
	matches := re.FindStringSubmatch(original)
	if len(matches) == 0 {
		return original
	}
	message := matches[1]
	decrypt, err := Decrypt(password, message)
	if err != nil {
		return message
	}
	return decrypt
}
