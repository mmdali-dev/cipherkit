package md5

import (
	Md5 "crypto/md5"
	"encoding/hex"
)

// تابع برای رمزنگاری با MD5
func TextToMd5(input string) string {
	hash := Md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
