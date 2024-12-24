package rsa

import (
	"crypto/rand"
	Rsa "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// GenerateKeys generates a new RSA private and public key pair.
func GenerateKeys(bits int) (*Rsa.PrivateKey, *Rsa.PublicKey, error) {
	privateKey, err := Rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func PublicKeyToString(pub *Rsa.PublicKey) (string, error) {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pub)
	pemEncodedPubKey := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubKeyBytes})
	return string(pemEncodedPubKey), nil
}

func PrivateKeyToString(priv *Rsa.PrivateKey) (string, error) {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(priv)
	pemEncodedPrivKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privKeyBytes})
	return string(pemEncodedPrivKey), nil
}

func StringToPublicKey(s string) (*Rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func StringToPrivateKey(s string) (*Rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
