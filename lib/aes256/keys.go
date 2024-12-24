package aes256

import (
	"crypto/rand"
	"errors"

	"github.com/mmdali-dev/cipherkit/lib/base64"
)

// GenerateKey creates a new AES-256 key
func GenerateKey() ([]byte, error) {
	var key [32]byte
	_, err := rand.Read(key[:])
	if err != nil {
		return nil, err
	}
	return key[:], nil
}

// KeyToString converts a Key to a string
func KeyToString(key []byte) string {
	return base64.Encode(key)
}

// StringToKey converts a string to a Key
func StringToKey(s string) ([]byte, error) {

	key, err := base64.Decode(s)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes")
	}
	return key, nil
}
