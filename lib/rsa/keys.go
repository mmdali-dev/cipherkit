package rsa

import (
	"crypto/rand"
	Rsa "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GeneratePrivateKey(bits int) (*Rsa.PrivateKey, error) {
	return Rsa.GenerateKey(rand.Reader, bits)
}

func GeneratePublicKey(privateKey *Rsa.PrivateKey) *Rsa.PublicKey {
	return &privateKey.PublicKey
}

func PrivateKeyToString(privateKey *Rsa.PrivateKey) string {
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	return string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}))
}

func PublicKeyToString(publicKey *Rsa.PublicKey) string {
	pubBytes := x509.MarshalPKCS1PublicKey(publicKey)
	return string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubBytes}))
}

func StringToPublicKey(publicKeyStr string) (*Rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	return pub, err
}

func StringToPrivateKey(privateKeyStr string) (*Rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv, err
}
