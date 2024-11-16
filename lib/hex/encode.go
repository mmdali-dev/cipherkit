package hex

import (
	Hex "encoding/hex"
)

// تابع انکود
func TextToHex(input string) string {
	encoded := Hex.EncodeToString([]byte(input))
	return encoded
}
