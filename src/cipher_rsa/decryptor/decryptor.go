package main

import "fmt"
import "flag"
import "cipher_rsa/common"

func decrypt(msg int) int {
	var decrypted int
	var publicKey, privateKey []int

	publicKey, privateKey = common.GetPublicAndPrivateKeys()
	fmt.Println(publicKey, privateKey)

	var n, d int = privateKey[0], privateKey[1]

	decrypted = common.PowMod(msg, d, n)

	return decrypted
}

func main() {
	n := flag.Int("n", 0, "source number")
	flag.Parse()

	fmt.Println("Source:", *n)
	decrypted := decrypt(*n)
	fmt.Println("Decrypted:", decrypted)
}
