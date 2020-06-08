package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func pkcs5Padding(content []byte, blockSize int) []byte {
	padding := blockSize - len(content)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(content, padText...)
}

func pkcs5UnPadding(cipherText []byte) []byte {
	length := len(cipherText)
	unPadding := int(cipherText[length-1])
	return cipherText[:(length - unPadding)]
}

// aes_256_cbc encrypt
func AesEncrypt(content []byte, key []byte) ([]byte, error) {
	keyBuff := make([]byte, 32)
	copy(keyBuff, key)
	block, err := aes.NewCipher(keyBuff)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	content = pkcs5Padding(content, blockSize)
	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(content))
	blockMode.CryptBlocks(encrypted, content)
	result := make([]byte, 16+len(encrypted))
	copy(result, iv)
	copy(result[16:], encrypted)
	return result, nil
}

// aes_256_cbc decrypt
func AesDecrypt(cipherText []byte, key []byte) ([]byte, error) {
	keyBuff := make([]byte, 32)
	copy(keyBuff, key)
	block, err := aes.NewCipher(keyBuff)
	if err != nil {
		return nil, err
	}
	iv := cipherText[0:16]
	cipherText = cipherText[16:]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(cipherText))
	blockMode.CryptBlocks(decrypted, cipherText)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted, nil
}

// encrypt (aes_256_cbc) for url safe
func EncryptString(content, key string) (string, error) {
	b, err := AesEncrypt([]byte(content), []byte(key))
	if err != nil {
		return "", err
	}
	result := base64.RawURLEncoding.EncodeToString(b)
	return result, nil
}

// decrypt (aes_256_cbc) for url safe
func DecryptString(cipherText, key string) (string, error) {
	b, err := base64.RawURLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	result, err := AesDecrypt(b, []byte(key))
	if err != nil {
		return "", err
	}
	return string(result), nil
}
