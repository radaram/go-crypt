package main

import (
	"fmt"
	"flag"
	"cipher_trithemius/common"
)


func getDecryptedLetterCode(alphabetPos int, textPos int, alphabetSize int) int {
	return (alphabetPos - common.GetOffsetStep(textPos)) % alphabetSize
}


func decrypt(text string, alphabet map[string]int) string {
	alphabetSize := len(alphabet)
	alphabetReversed := common.ReverseMap(alphabet)
	var decrypted string
	for i, ch := range text {
		code := getDecryptedLetterCode(alphabet[string(ch)], i, alphabetSize)
		decrypted += alphabetReversed[code]
	}

	return decrypted
}

func main() {
    text := flag.String("text", "", "encrypted text")
	flag.Parse()
	
	fmt.Println("Source:", *text)	
	
	decrypted := decrypt(*text, common.Alphabet)
	fmt.Println("decrypted:", decrypted)
}
