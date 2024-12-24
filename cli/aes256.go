package cli

import (
	"fmt"
	"strings"

	"github.com/mmdali-dev/cipherkit/internal"
	"github.com/mmdali-dev/cipherkit/lib/aes256"
)

func Aes256Handler(gen, dec bool, text, skey, file, output, keyfile string) {

	if gen {
		ky, err := aes256.GenerateKey()
		if err != nil {
			panic("error generating key: " + err.Error())
		}
		aes256.SaveKey("key.txt", aes256.KeyToString(ky))
		fmt.Println("key created at : key.txt!")
	} else {
		//validate key from cli or path.
		var key string
		var err error
		if skey == "" {
			if keyfile == "" {
				keyfile = "key.txt"
			}
			key, err = aes256.LoadKey(keyfile)
			if err != nil {
				panic("error load key from keyfile. please set --keyfile=path or --key=key and run.")
			}
		} else {
			key = skey
		}

		if text != "" {
			if !dec { //encrypt text
				result, err := aes256.Encrypt(key, []byte(text))
				if err != nil {
					panic("error encrypt text: " + err.Error())
				}
				fmt.Println(result)
			} else { //decrypt text
				result, err := aes256.Decrypt(key, text)
				if err != nil {
					panic("error Decrypt data: " + err.Error())
				}
				fmt.Println(string(result))
			}
		} else { //file method
			if file == "" {
				panic("error , please enter the --file path or --text for message")
			}
			if output == "" {
				output = file
			}
			data, err := internal.ReadFromFile(file)
			if err != nil {
				panic("error reading file: " + err.Error())
			}
			if !dec { //encrypt file
				result, err := aes256.Encrypt(key, data)
				if err != nil {
					panic("error encrypt data: " + err.Error())
				}
				err = internal.WriteToFile(output, []byte(result))
				if err != nil {
					panic("error writing file: " + err.Error())
				}
			} else { //decrypt file
				result, err := aes256.Decrypt(key, string(data))
				if err != nil {
					panic("error decrypt data: " + err.Error())
				}
				if strings.HasSuffix(file, ".txt") || strings.HasSuffix(output, ".txt") {
					err = internal.WriteToFile(output, []byte(string(result)))
				} else {
					err = internal.WriteToFile(output, result)
				}
				if err != nil {
					panic("error writing file: " + err.Error())
				}
			}
		}
	}
}
