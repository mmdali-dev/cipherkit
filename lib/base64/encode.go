package base64

import (
	Base64 "encoding/base64"
)

func TextToBase64(input string) string {
	return Base64.StdEncoding.EncodeToString([]byte(input))
}
