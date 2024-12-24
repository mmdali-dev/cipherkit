package binary

import (
	"fmt"
	"strings"
)

func Encode(text string) string {
	var binary strings.Builder
	for _, char := range text {
		binary.WriteString(fmt.Sprintf("%08b ", char))
	}
	return strings.TrimSpace(binary.String())
}
