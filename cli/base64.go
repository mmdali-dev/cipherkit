package cli

import (
	"fmt"
	"strings"

	"github.com/mmdali-dev/cipherkit/internal"
	"github.com/mmdali-dev/cipherkit/lib/base64"
)

func Base64Handler(text, file, output string, dec bool) {
	//text mode
	if text != "" { //encode data
		if !dec {
			fmt.Println(base64.Encode([]byte(text)))
		} else { //decode data
			decdata, err := base64.Decode(text)
			if err != nil {
				panic("error decoding data : " + err.Error())
			}
			fmt.Println(string(decdata))
		}
		//file mode
	} else if file != "" {
		//check output is exist
		if output == "" {
			output = file
		}
		//reading data from file
		data, err := internal.ReadFromFile(file)
		if err != nil {
			panic("error reading file: " + err.Error())
		}
		//check file size
		if len(data) < 1 {
			panic("file is empty.")
		}
		//encode file
		if !dec {
			encdata := base64.Encode(data)
			internal.WriteToFile(output, []byte(encdata))
		} else { //decode file
			decdata, err := base64.Decode(string(data))
			if err != nil {
				panic("error decoding data : " + err.Error())
			}
			//check normal characters file / .txt
			if strings.HasSuffix(file, ".txt") {
				internal.WriteToFile(output, []byte(string(decdata)))
			} else {
				internal.WriteToFile(output, decdata)
			}
		}

	}
}
