package main

import (
	"fmt"
	"flag"
	"cipher_trithemius/common"
)


func getEncryptedLetterCode(alphabetPos int, textPos int, alphabetSize int) int {
	return (alphabetPos + common.GetOffsetStep(textPos)) % alphabetSize
}


func encrypt(text string, alphabet map[string]int) string {
	alphabetSize := len(alphabet)
	alphabetReversed := common.ReverseMap(alphabet)
	var encrypted string
	for i, ch := range text {
		code := getEncryptedLetterCode(alphabet[string(ch)], i, alphabetSize)
		encrypted += alphabetReversed[code]
	}

	return encrypted
}

func main() {
    text := flag.String("text", "", "source text")
	flag.Parse()
	
	fmt.Println("Source:", *text)	
	
	encrypted := encrypt(*text, common.Alphabet)
	fmt.Println("Encrypted:", encrypted)
}
