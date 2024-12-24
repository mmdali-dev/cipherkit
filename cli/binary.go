package cli

import (
	"fmt"

	"github.com/mmdali-dev/cipherkit/lib/binary"
)

func BinaryHandler(text string, dec bool) {
	if text != "" {
		if !dec {
			fmt.Println(binary.Encode(text))
		} else {
			data, err := binary.Decode(text)
			if err != nil {
				panic("error decoding text: " + err.Error())
			}
			fmt.Println(data)
		}
	} else {
		panic("error , please use --text to input")
	}
}
