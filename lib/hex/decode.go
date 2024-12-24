package hex

import Hex "encoding/hex"

func Decode(input string) ([]byte, error) {
	decoded, err := Hex.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
