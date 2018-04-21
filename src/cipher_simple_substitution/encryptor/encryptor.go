package main

import "fmt"
import "flag"

import "cipher_simple_substitution/common"

func main() {
	text := flag.String("text", "", "source text")
	flag.Parse()

	fmt.Println("Source:", *text)
	encrypted := common.Crypt(*text, common.SourceAlphabet, common.EncryptAlphabet)
	fmt.Println("Encrypted:", encrypted)
}
