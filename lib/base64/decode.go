package base64

import (
	Base64 "encoding/base64"
)

func Decode(input string) ([]byte, error) {
	decodedBytes, err := Base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return decodedBytes, nil
}
