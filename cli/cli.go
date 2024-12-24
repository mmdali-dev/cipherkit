package cli

import (
	"flag"
	"fmt"
	// Adjust the import path as needed
)

func Cli() {
	// Define command-line flags
	a := flag.Bool("a", false, "show example usage")
	mod := flag.String("mod", "", "sha256 , md5 , hex , binary , rsa  , aes256")
	text := flag.String("text", "", "use for input inline text like : --text=\"hello cipherkit\". or encoded text to decrypt")
	file := flag.String("file", "", "use for target file path you want to decrypt/encrypt. --file=filepath")
	output := flag.String("output", "", "use for result file path you want to decrypt/encrypt. --output=filepath")
	dec := flag.Bool("dec", false, "use when you want to decrypt text or file. uage : -dec (empty)")
	gen := flag.Bool("gen", false, "use when you want to generate keyfile. uage : -gen (empty). in rsa and aes256")
	keyfile := flag.String("keyfile", "", "key path to find and set in rsa and aes256.")
	key := flag.String("key", "", "set inline key when you dont want use file. like : -key=somekey ")
	flag.Parse()

	if *a {
		fmt.Println("cipherkit uasge example for all methods.")
		fmt.Println("cipherkit -mod rsa -gen //generate keys")

		fmt.Println("cipherkit -mod rsa -text \"hello guys\" -key public.pem(optional) //encrypt inline text")
		fmt.Println("cipherkit -mod rsa -text \"encoded text\" -key private.pem(optional) -dec //decrypt inline text")

		fmt.Println("cipherkit -mod rsa -file filepath -key public.pem(optional) -output out.enc(optional) //encrypt file")
		fmt.Println("cipherkit -mod rsa -file filepath -key private.pem(optional) -dec -output out.enc(optional) //decrypt file")

		fmt.Println("\n\ncipherkit -mod aes256 -gen //generate key")

		fmt.Println("cipherkit -mod aes256 -text \"hello guys\" -keyfile key.txt(optional) //encrypt inline text")
		fmt.Println("cipherkit -mod aes256 -text \"hello guys\" -key yourkey(optional) //encrypt inline text with inline key")
		fmt.Println("cipherkit -mod aes256 -text \"encoded text\" -keyfile key.txt(optional) -dec //decrypt inline text")
		fmt.Println("cipherkit -mod aes256 -text \"encoded text\" -key yourkey(optional) -dec //decrypt inline text with inline key")

		fmt.Println("cipherkit -mod aes256 -file filepath -keyfile key.txt(optional) -output out.enc(optional) //encrypt file")
		fmt.Println("cipherkit -mod aes256 -file filepath -keyfile key.txt(optional) -dec -output out.enc(optional) //decrypt file")

		fmt.Println("cipherkit -mod aes256 -file filepath -key yourkey(optional) -output out.enc(optional) //encrypt file with inline key")
		fmt.Println("cipherkit -mod aes256 -file filepath -key yourkey(optional) -dec -output out.enc(optional) //decrypt file with inline key")

		fmt.Println("\n\nhere is one example and you can replace base64 with this methods[hex, binary, md5, sha256, ]:")
		fmt.Println("cipherkit -mod base64 -text \"your message\"")
		fmt.Println("cipherkit -mod base64 -text \"encoded message\" -dec")

		fmt.Println("cipherkit -mod base64 -file filepath -output out.enc(optional)")
		fmt.Println("cipherkit -mod base64 -file filepath -output out.enc(optional) -dec")
	}
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
