package jasypt

import (
	"testing"
)

func TestBasicTextEncryptor(t *testing.T) {
	password := "awesomeTang"
	ciphertext1 := "ENC(zH0ANx/qIJ4HDNd2SrDOig==)"
	ciphertext2 := "123456"

	enc := CreateBasicTextEncryptor(password)
	plainText1 := enc.GetString(ciphertext1)
	if plainText1 != "123456" {
		t.Fail()
	}

	plainText2 := enc.GetString(ciphertext2)
	if plainText2 != "123456" {
		t.Fail()
	}
}
