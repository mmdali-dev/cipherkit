package md5

import (
	Md5 "crypto/md5"
	"encoding/hex"
)

// تابع برای رمزنگاری با MD5
func Encode(input []byte) string {
	hash := Md5.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}
