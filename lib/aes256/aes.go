package aes256

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"errors"

	"github.com/mmdali-dev/cipherkit/lib/base64"
)

// Encrypt encrypts plaintext using the given key
func Encrypt(skey string, plaintext []byte) (string, error) {
	key, err := StringToKey(skey)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return base64.Encode(ciphertext), nil
}

// Decrypt decrypts ciphertext using the given key
func Decrypt(skey, encodedtext string) ([]byte, error) {
	key, err := StringToKey(skey)
	if err != nil {
		return nil, err
	}
	data, err := base64.Decode(encodedtext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
