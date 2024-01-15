package encoders

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func EncryptPayload(payload string, key []byte) (string, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(payload), nil)

	encrypted := append(nonce, ciphertext...)

	return hex.EncodeToString(encrypted), nil
}

func DecryptPayload(encryptedText string, key []byte) (string, error) {

	encrypted, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return "", errors.New("text to short")
	}

	nonce, cipherTXT := encrypted[:nonceSize], encrypted[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherTXT, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
