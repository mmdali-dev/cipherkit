package hex

import (
	Hex "encoding/hex"
)

// تابع انکود
func TextToHex(input []byte) string {
	encoded := Hex.EncodeToString(input)
	return encoded
}
