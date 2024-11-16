package hex

import Hex "encoding/hex"

func HexToText(input string) (string, error) {
	decoded, err := Hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
