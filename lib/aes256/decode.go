package aes256

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, errors.New("padding size error")
	}
	return src[:(length - unpadding)], nil
}

func Decrypt(ciphertext, key string) (string, error) {
	ciphertextWithIV, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertextWithIV) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertextWithIV[:aes.BlockSize]
	ciphertextWithoutIV := ciphertextWithIV[aes.BlockSize:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	paddedText := make([]byte, len(ciphertextWithoutIV))
	mode.CryptBlocks(paddedText, ciphertextWithoutIV)

	plaintext, err := unpad(paddedText)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
