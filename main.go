package main

import (
	"fmt"

	"github.com/mmdali-dev/cipherkit/lib/aes256"
)

func main() {
	key, _ := aes256.LoadKey("a")
	fmt.Println(key)
	bkey, _ := aes256.StringToKey(key)
	encoded, _ := aes256.Encrypt(bkey, []byte("mohammad"))
	fmt.Println(encoded)
	decoded, _ := aes256.Decrypt(bkey, encoded)
	fmt.Println(string(decoded))
}
