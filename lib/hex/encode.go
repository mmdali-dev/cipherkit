package hex

import (
	Hex "encoding/hex"
)

// تابع انکود
func Encode(input []byte) string {
	encoded := Hex.EncodeToString(input)
	return encoded
}
