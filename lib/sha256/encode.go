package sha256

import (
	Sha256 "crypto/sha256"
	"encoding/hex"
)

// تابع انکد
func Encode(input []byte) string {
	hash := Sha256.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}
