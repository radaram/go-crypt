package common


func Decrypt(text string, shift int, size int) string {
	var decrypted string
	var old_code, new_code int
	for  _, ch := range text {
		old_code = int(ch)
		new_code = old_code - shift
		if ch >= 'a' && ch <= 'z' && new_code < int('a') || 
		   ch >= 'A' && ch <= 'Z' && new_code < int('A') {
			new_code += size
		}
		
		decrypted += string(new_code)
	}
	return decrypted
}

