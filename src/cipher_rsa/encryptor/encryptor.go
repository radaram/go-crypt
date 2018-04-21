package main

import "fmt"
import "flag"
import "cipher_rsa/common"

func encrypt(msg int) int {
	var encrypted int
	var publicKey, privateKey []int

	publicKey, privateKey = common.GetPublicAndPrivateKeys()
	fmt.Println(publicKey, privateKey)

	var n, e int = publicKey[0], publicKey[1]

	encrypted = common.PowMod(msg, e, n)

	return encrypted
}

func main() {
	n := flag.Int("n", 0, "source number")
	flag.Parse()

	fmt.Println("Source:", *n)
	encrypted := encrypt(*n)
	fmt.Println("Encrypted:", encrypted)
}
