package base64

import (
	Base64 "encoding/base64"
)

func Encode(input []byte) string {
	return Base64.StdEncoding.EncodeToString(input)
}
