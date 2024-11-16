package rsa

import (
	"crypto/rand"
	Rsa "crypto/rsa"
)

func RsaToText(privateKey *Rsa.PrivateKey, ciphertext []byte) (string, error) {
	plaintext, err := Rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	return string(plaintext), err
}
