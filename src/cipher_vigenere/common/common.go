package common

func Upper(text string) string {
	var newText string
	var offset int
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			offset = 32
		} else if ch >= 'A' && ch <= 'Z' {
			offset = 0
		}

		newText += string(int(ch) - offset)
	}

	return newText
}
