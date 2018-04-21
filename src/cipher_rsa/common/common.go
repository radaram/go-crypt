package common

import "crypto/rand"
import "log"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func isFloat(n float64) bool {
	return n == float64(int64(n))
}

func GetPublicAndPrivateKeys() ([]int, []int) {
	pBig, err := rand.Prime(rand.Reader, 8)
	if err != nil {
		panic(err)
	}

	qBig, err := rand.Prime(rand.Reader, 8)
	if err != nil {
		panic(err)
	}

	p, q := int(pBig.Int64()), int(qBig.Int64())
	p, q = 29, 31
	n := p * q
	f_n := (p - 1) * (q - 1) // function of Euler

	var d float64 = 0
	var e int = 1
	var attempts = 0

	for d == 0 || !isFloat(d) {
		if attempts > 100 {
			log.Fatalf("Unable to retrieve secret key")
		}

		attempts += 1
		e += 2

		for gcd(f_n, e) != 1 {
			e += 2
		}

		for k := 1; k <= 10 && (d == 0 || !isFloat(d)); k++ {
			d = float64(k*f_n+1) / float64(e)
		}
	}
	return []int{n, e}, []int{n, int(d)}
}

func PowMod(a, b, m int) int {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}
