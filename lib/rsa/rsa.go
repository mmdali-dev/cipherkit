package rsa

import (
	"crypto/rand"
	Rsa "crypto/rsa"
	"crypto/sha256"

	"github.com/mmdali-dev/cipherkit/lib/base64"
)

// Encrypt encrypts the plaintext using the provided public key.
func Encrypt(publicKey *Rsa.PublicKey, plaintext []byte) (string, error) {

	hashed := sha256.New()
	encodedtext, err := Rsa.EncryptOAEP(hashed, rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return "", err
	}
	return base64.Encode(encodedtext), nil

}

// Decrypt decrypts the ciphertext using the provided private key.
func Decrypt(privateKey *Rsa.PrivateKey, encodedtext string) ([]byte, error) {
	ciphertext, err := base64.Decode(encodedtext)
	if err != nil {
		return nil, err
	}
	hashed := sha256.New()
	return Rsa.DecryptOAEP(hashed, rand.Reader, privateKey, ciphertext, nil)
}
