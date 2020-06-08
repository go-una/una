package tools

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	key := []byte("abcdefg")
	content := []byte("ABC你好")
	_, err := AesEncrypt(content, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAesEncrypt2(t *testing.T) {
	key := []byte("abcdefg")
	var content []byte
	_, err := AesEncrypt(content, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAesDecrypt(t *testing.T) {
	key := []byte("abcdefgxxxxxx")
	content := []byte("ABC你好")
	cipherText, err := AesEncrypt(content, key)
	if err != nil {
		t.Fatal(err)
	}
	result, err := AesDecrypt(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(result, content) != 0 {
		t.Fatalf("AesDecrypt result is invalid")
	}
}

func TestEncryptString(t *testing.T) {
	key := "secret-key"
	content := "说好的加密呢😁 abc\n123"
	cipherText, err := EncryptString(content, key)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ContainsAny(cipherText, "+/=") {
		t.Fatal(errors.New("EncryptString result is not url safe"))
	}
}

func TestDecryptString(t *testing.T) {
	key := "secret-key"
	content := "说好的加密呢😁 abc\n123"
	cipherText, err := EncryptString(content, key)
	if err != nil {
		t.Fatal(err)
	}
	result, err := DecryptString(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}
	if result != content {
		t.Fatal(errors.New("DecryptString result is invalid"))
	}
}
