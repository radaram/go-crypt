package main

import "fmt"
import "flag"

import "cipher_vigenere/common"

func decrypt(text string, key string, size int) string {
	var decrypted string
	var chKey byte
	text, key = common.Upper(text), common.Upper(key)

	for i, ch := range text {
		chKey = key[i%len(key)]
		decrypted += string(((((int(ch) - 'A') - (int(chKey) - 'A')) + size) % size) + 'A')
	}

	return decrypted

}

func main() {
	text := flag.String("text", "", "enctypted text")
	key := flag.String("key", "", "key")
	flag.Parse()

	size := 26

	fmt.Println("Encrypted:", *text)
	decrypted := decrypt(*text, *key, size)
	fmt.Println("Decrypted:", decrypted)
}
