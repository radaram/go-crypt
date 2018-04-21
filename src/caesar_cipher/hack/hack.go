package main

import "fmt"
import "flag"

import "caesar_cipher/common"


func hack(text string, size int) {
	var decrypted string
	for i:=1; i <= size; i++ {
		decrypted = common.Decrypt(text, i, size)
		fmt.Printf("shift %d: %s\n", i, decrypted)
	}
}


func main() {
	text := flag.String("text", "", "encrypted text")
	flag.Parse()
	
	fmt.Println("Encrypted:", *text)	
	size := 26
	hack(*text, size)
}
