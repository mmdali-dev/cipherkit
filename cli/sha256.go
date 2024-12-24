package cli

import (
	"fmt"

	"github.com/mmdali-dev/cipherkit/internal"
	"github.com/mmdali-dev/cipherkit/lib/sha256"
)

func Sha256Handler(text, file, output string) {
	if text != "" {
		fmt.Println(sha256.Encode([]byte(text)))
	} else if file != "" {
		data, err := internal.ReadFromFile(file)
		if err != nil {
			panic("error reading file: " + err.Error())
		}
		if len(data) < 1 {
			panic("file is empty.")
		}
		encdata := sha256.Encode(data)
		if output == "" {
			output = file
		}
		internal.WriteToFile(output, []byte(encdata))
	}
}
