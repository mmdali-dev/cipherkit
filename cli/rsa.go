package cli

import (
	"fmt"
	"strings"

	"github.com/mmdali-dev/cipherkit/internal"
	"github.com/mmdali-dev/cipherkit/lib/rsa"
)

func RsaHandler(gen, dec bool, text, key, file, output string) {
	if gen {
		pri, pub, err := rsa.GenerateKeys(2048)
		if err != nil {
			panic("error generating keys: " + err.Error())
		}
		rsa.SavePrivateKey("private.pem", pri)
		rsa.SavePublicKey("public.pem", pub)
		fmt.Println("keys created at : private.pem , public.pem!")
	} else {
		if key == "" {
			if dec {
				key = "private.pem"
			} else {
				key = "public.pem"
			}
		}
		if text != "" {
			if !dec {
				pub, err := rsa.LoadPublicKey(key)
				if err != nil {
					panic("error load public key: " + err.Error())
				}
				result, err := rsa.Encrypt(pub, []byte(text))
				if err != nil {
					panic("error encrypt text: " + err.Error())
				}
				fmt.Println(result)
			} else {
				pri, err := rsa.LoadPrivateKey(key)
				if err != nil {
					panic("error load private key: " + err.Error())
				}
				result, err := rsa.Decrypt(pri, text)
				if err != nil {
					panic("error Decrypt data: " + err.Error())
				}
				fmt.Println(string(result))
			}
		} else {
			if file == "" {
				panic("error , please enter the --file path")
			}
			if output == "" {
				output = file
			}
			data, err := internal.ReadFromFile(file)
			if err != nil {
				panic("error reading file: " + err.Error())
			}
			if !dec {
				pub, err := rsa.LoadPublicKey(key)
				if err != nil {
					panic("error load public key: " + err.Error())
				}
				result, err := rsa.Encrypt(pub, data)
				if err != nil {
					panic("error encrypt data: " + err.Error())
				}
				err = internal.WriteToFile(output, []byte(result))
				if err != nil {
					panic("error writing file: " + err.Error())
				}
			} else {
				pri, err := rsa.LoadPrivateKey(key)
				if err != nil {
					panic("error load private key: " + err.Error())
				}
				result, err := rsa.Decrypt(pri, string(data))
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
