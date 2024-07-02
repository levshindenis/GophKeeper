package tools

import (
	"crypto/aes"
	"crypto/cipher"
	rb "crypto/rand"
	"encoding/base64"
)

// GenerateCrypto - создает крипто-ключ
func GenerateCrypto(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rb.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateCookie создает куки пользователя.
func GenerateCookie(value string) (string, error) {
	key, err := GenerateCrypto(aes.BlockSize)
	if err != nil {
		return "", err
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	nonce, err := GenerateCrypto(aesgcm.NonceSize())
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(
		aesgcm.Seal(nil, nonce, []byte(value), nil)), nil
}
