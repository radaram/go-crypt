package main

import "fmt"
import "flag"
import "log"
import "strconv"

func encrypt(text string, key string, size int) string {
	key_index := 0
	key_len := len(key)
	var encrypted string
	var old_code, new_code int

	for _, ch := range text {
		shift, err := strconv.Atoi(string(key[key_index]))
		if err != nil {
			log.Fatalf("Key must be number. %s", err)
		}
		old_code = int(ch)
		new_code = old_code + shift

		if ch >= 'a' && ch <= 'z' && new_code > int('z') ||
			ch >= 'A' && ch <= 'Z' && new_code > int('Z') {
			new_code -= size
		}

		encrypted += string(new_code)

		key_index++
		if key_index >= key_len {
			key_index = 0
		}
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
