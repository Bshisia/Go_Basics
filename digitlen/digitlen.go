package main

import "fmt"

func DigitLen(n, base int) int {
	if n < 0 {
		n = -n
	}

	if base < 2 || base > 36 {
		return -1
	}

	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))  // Output: 3
	fmt.Println(DigitLen(100, 2))   // Output: 7
	fmt.Println(DigitLen(-100, 16)) // Output: 2
	fmt.Println(DigitLen(100, -1))  // Output: -1
}
