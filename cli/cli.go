package cli

import (
	"flag"
	"fmt"
	// Adjust the import path as needed
)

func Cli() {
	// Define command-line flags
	mod := flag.String("mod", "", "sha256 --text=hi, ")
	text := flag.String("text", "", "--text=\"hello cipherkit\"")
	file := flag.String("file", "", "--file=filepath ")
	output := flag.String("output", "", "--output=filepath")
	dec := flag.Bool("dec", false, "use for decrypt in algorythms.")

	flag.Parse()
	switch *mod {
	case "sha256":
		Sha256Handler(*text, *file, *output)
	case "md5":
		Md5Handler(*text, *file, *output)
	case "base64":
		Base64Handler(*text, *file, *output, *dec)
	default:
		fmt.Println("use --help to see how uuse the cipherkit")
	}
}
