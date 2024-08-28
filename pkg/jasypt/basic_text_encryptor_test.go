package jasypt

import (
	"testing"
)

func TestBasicTextEncryptor(t *testing.T) {
	password := "awesomeTang"
	ciphertext1 := "ENC(zH0ANx/qIJ4HDNd2SrDOig==)"
	ciphertext2 := "123456"

	plainText1 := GetString(password, ciphertext1)
	if plainText1 != "123456" {
		t.Fail()
	}

	plainText2 := GetString(password, ciphertext2)
	if plainText2 != "123456" {
		t.Fail()
	}
}
