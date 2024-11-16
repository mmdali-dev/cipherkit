package rsa

import (
	"crypto/rand"
	Rsa "crypto/rsa"
)

func TextToRsa(publicKey *Rsa.PublicKey, plaintext string) ([]byte, error) {
	return Rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plaintext))
}
