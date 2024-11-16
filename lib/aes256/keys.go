package aes256

import "crypto/sha256"

func GenerateKey(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}
