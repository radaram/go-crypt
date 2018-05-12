package main

import "math/big"
import "crypto/rand"
import "fmt"
import "log"
import "sync"


var wg sync.WaitGroup


func genRandPrimeNum(optional ...int) *big.Int {
	var bits int = 8
    if len(optional) > 1 {
        panic("Too much args")
    } else if len(optional) == 1 {
        bits = optional[0]
    }
	
	pBig, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	return pBig
}


func genRandNaturalNumber(optional ...int64) *big.Int {
	var max int64 = 10000
    if len(optional) > 1 {
        panic("Too much args")
    } else if len(optional) == 1 {
        max = optional[0]
    }

	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}

	return nBig 
}



func PowMod(a, b, m int64) int64 {
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


func bob(in chan []int64, out chan []int64) {
	defer wg.Done()
	msg := <- in
	
	b := genRandNaturalNumber().Int64()
	g, p, A := msg[0], msg[1], msg[2]
	
	B := PowMod(g, b, p)	
	fmt.Println("Bob public key:", B)

	out <- []int64{B}
	
	K := PowMod(A, b, p)
	fmt.Println("General secret key:", K)
}


func alice(in chan []int64, out chan []int64) {
    defer wg.Done()
	p := genRandPrimeNum(16).Int64()
	g := genRandPrimeNum().Int64()
	max_attempts := 100
	
	for attempts := 1; PowMod(g, p-1, p) != 1; attempts++ {
		if attempts > max_attempts {
			log.Fatalf("Exceeded number of attempts")
		}
        g = genRandPrimeNum().Int64() 
	}  

	a := genRandNaturalNumber().Int64()
	A := PowMod(g, a, p)	
	fmt.Println("Alice public key:", A)

	out <- []int64{g, p, A}

	msg := <- in
	B := msg[0]
	
	K := PowMod(B, a, p)
	fmt.Println("General secret key:", K)
}

func main() {
	var aliceCh chan []int64 = make(chan []int64)
    var bobCh chan []int64 = make(chan []int64)
	
	wg.Add(1)
	go bob(bobCh, aliceCh)
	
	wg.Add(1)
	go alice(aliceCh, bobCh)

	wg.Wait()
}
