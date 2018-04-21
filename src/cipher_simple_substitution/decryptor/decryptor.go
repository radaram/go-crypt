package main

import "fmt"
import "flag"

import "cipher_simple_substitution/common"

func main() {
	text := flag.String("text", "", "encrypted text")
	flag.Parse()

	fmt.Println("Encrypted:", *text)
	decrypted := common.Crypt(*text, common.EncryptAlphabet, common.SourceAlphabet)
	fmt.Println("Decrypted:", decrypted)
}
