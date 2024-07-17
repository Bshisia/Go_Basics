package main

import "fmt"

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	} 
	if nb % 2 == 0 || nb % 3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for nb > 0 {
		if IsPrime(nb) {
			return nb
		}
		nb --
	}
	return nb
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}