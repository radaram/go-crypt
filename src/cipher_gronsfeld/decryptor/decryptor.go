package main

import "fmt"
import "flag"
import "log"
import "strconv"

func decrypt(text string, key string, size int) string {
	key_index := 0
	key_len := len(key)
	var decrypted string
	var old_code, new_code int

	for _, ch := range text {
		shift, err := strconv.Atoi(string(key[key_index]))
		if err != nil {
			log.Fatalf("Key must be number. %s", err)
		}
		old_code = int(ch)
		new_code = old_code - shift

		if ch >= 'a' && ch <= 'z' && new_code < int('a') ||
			ch >= 'A' && ch <= 'Z' && new_code < int('A') {
			new_code += size
		}

		decrypted += string(new_code)

		key_index++
		if key_index >= key_len {
			key_index = 0
		}
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
