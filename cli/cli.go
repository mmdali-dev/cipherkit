package cli

import (
	"flag"
	"fmt"
	// Adjust the import path as needed
)

func Cli() {
	// Define command-line flags
	mod := flag.String("mod", "", "sha256 , md5 , hex , binary , rsa  , aes256")
	text := flag.String("text", "", "use for input inline text like : --text=\"hello cipherkit\". or encoded text to decrypt")
	file := flag.String("file", "", "use for target file path you want to decrypt/encrypt. --file=filepath")
	output := flag.String("output", "", "use for result file path you want to decrypt/encrypt. --output=filepath")
	dec := flag.Bool("dec", false, "use when you want to decrypt text or file. uage : -dec (empty)")
	gen := flag.Bool("gen", false, "use when you want to generate keyfile. uage : -gen (empty). in rsa and aes256")
	keyfile := flag.String("keyfile", "", "key path to find and set in rsa and aes256.")
	key := flag.String("key", "", "set inline key when you dont want use file. like : -key=somekey ")
	flag.Parse()
	switch *mod {
	case "sha256":
		Sha256Handler(*text, *file, *output)
	case "md5":
		Md5Handler(*text, *file, *output)
	case "base64":
		Base64Handler(*text, *file, *output, *dec)
	case "hex":
		HexHandler(*text, *file, *output, *dec)
	case "binary":
		BinaryHandler(*text, *dec)
	case "rsa":
		RsaHandler(*gen, *dec, *text, *keyfile, *file, *output)
	case "aes256":
		Aes256Handler(*gen, *dec, *text, *key, *file, *output, *keyfile)
	default:
		fmt.Println("use --help to see how uuse the cipherkit")
	}
}
