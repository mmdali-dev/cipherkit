package binary

import (
	"fmt"
	"strings"
)

func BinaryToText(binary string) (string, error) {
	var text strings.Builder
	binaryValues := strings.Fields(binary)

	for _, bin := range binaryValues {
		var char byte
		_, err := fmt.Sscanf(bin, "%b", &char)
		if err != nil {
			return "", err
		}
		text.WriteByte(char)
	}

	return text.String(), nil
}
