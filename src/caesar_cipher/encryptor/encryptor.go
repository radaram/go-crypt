package main

import "fmt"
import "flag"
import "log"


func encrypt(text string, shift int, size int) string {
	var encrypted string
	var old_code, new_code int
	for  _, ch := range text {
		old_code = int(ch)
		new_code = old_code + shift
		if ch >= 'a' && ch <= 'z' && new_code > int('z') || 
		   ch >= 'A' && ch <= 'Z' && new_code > int('Z') {
			new_code -= size
		}
		
		encrypted += string(new_code)
	}
	return encrypted
}


func main() {
	text := flag.String("text", "", "source text")
	shift := flag.Int("shift", 0, "shift")
	flag.Parse()
	
	size := 26
	if *shift > size {
		log.Fatalf("Cannot be more then %d", *shift)
	}
	fmt.Println("Source:", *text)	
	encrypted := encrypt(*text, *shift, size)
	fmt.Println("Encrypted":, encrypted)
}
