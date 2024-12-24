package rsa

import (
	Rsa "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// SavePublicKey saves the public key to a file.
func SavePublicKey(filename string, pub *Rsa.PublicKey) error {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pub)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = pem.Encode(f, &pem.Block{Type: "PUBLIC KEY", Bytes: pubKeyBytes})
	if err != nil {
		return err
	}
	return nil
}

// LoadPublicKey loads the public key from a file.
func LoadPublicKey(filename string) (*Rsa.PublicKey, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// SavePrivateKey saves the private key to a file.
func SavePrivateKey(filename string, priv *Rsa.PrivateKey) error {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(priv)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = pem.Encode(f, &pem.Block{Type: "PRIVATE KEY", Bytes: privKeyBytes})
	if err != nil {
		return err
	}
	return nil
}

// LoadPrivateKey loads the private key from a file.
func LoadPrivateKey(filename string) (*Rsa.PrivateKey, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
