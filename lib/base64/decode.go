package base64

import (
	Base64 "encoding/base64"
)

func Base64ToText(input string) (string, error) {
	decodedBytes, err := Base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
