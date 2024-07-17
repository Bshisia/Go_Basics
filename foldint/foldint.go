package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Itoa(nb int) string {
	result := ""
	for nb > 0 {
		digit := nb % 10
		result = string('0'+digit) + result
		nb /= 10
	}
	return result
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, slice := range a {
		n = f(n, slice)
	}
	result := Itoa(n)
	for _, res := range result {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}
