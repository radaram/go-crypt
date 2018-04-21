package common

import "strings"

const SourceAlphabet string = "abcdefghijklmnopqrstuvwxyz"
const EncryptAlphabet string = "iuqtrvzpbnfschdkljegxaymwo"

func Crypt(text string, alphabet1 string, alphabet2 string) string {
	var encrypted, letter string
	var i int

	for _, ch := range text {
		letter = string(ch)
		i = strings.Index(alphabet1, letter)
		if i == -1 {
			encrypted += letter
			continue
		}

		encrypted += string(alphabet2[i])
	}
	return encrypted
}
