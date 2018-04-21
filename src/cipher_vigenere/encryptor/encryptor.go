package main

import "fmt"
import "flag"

import "cipher_vigenere/common"

func encrypt(text string, key string, size int) string {
	var encrypted string
	var chKey byte
	text, key = common.Upper(text), common.Upper(key)

	for i, ch := range text {
		chKey = key[i%len(key)]
		encrypted += string((((int(ch) - 'A') + (int(chKey) - 'A')) % size) + 'A')
	}

	return encrypted

}

func main() {
	text := flag.String("text", "", "source text")
	key := flag.String("key", "", "key")
	flag.Parse()

	size := 26

	fmt.Println("Source:", *text)
	encrypted := encrypt(*text, *key, size)
	fmt.Println("Encrypted:", encrypted)
}
