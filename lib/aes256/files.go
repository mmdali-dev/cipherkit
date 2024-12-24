package aes256

import (
	"os"
)

// SaveKey stores a key in a file
func SaveKey(filename string, key string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(key))
	return err
}

// LoadKey retrieves a key from a file
func LoadKey(filename string) (string, error) {
	key, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(key), nil
}
