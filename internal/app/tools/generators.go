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
	key, err := GenerateCrypto(2 * aes.BlockSize)
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

func Encrypt(plaintext string, secretKey string) string {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonce, err := GenerateCrypto(gcm.NonceSize())
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(
		gcm.Seal(nonce, nonce, []byte(plaintext), nil))
}

func Decrypt(hashText string, secretKey string) string {
	cryptoBytes, _ := base64.StdEncoding.DecodeString(hashText)
	ciphertext := string(cryptoBytes)

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
