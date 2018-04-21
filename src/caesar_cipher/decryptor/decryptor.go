package main

import "fmt"
import "flag"
import "log"

import "caesar_cipher/common"


func main() {
	text := flag.String("text", "", "encrypted text")
	shift := flag.Int("shift", 0, "shift")
	flag.Parse()
	
	size := 26
	if *shift > size {
		log.Fatalf("Cannot be more then %d", *shift)
	}
	fmt.Println("Encrypted:", *text)	
	decrypted := common.Decrypt(*text, *shift, size)
	fmt.Println("Decrypted:", decrypted)
}
