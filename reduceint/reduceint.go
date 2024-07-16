package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	result := a[0]
	for _, num := range a[1:] {
		result = f(result, num)
	}
	str := Itoa(result)
	for _, val := range str {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		if cur == 0 {
			return acc // handle division by zero
		}
		return acc / cur
	}

	as := []int{500, 2}
	ReduceInt(as, mul) // Output: 1000 (500 * 2)
	ReduceInt(as, sum) // Output: 502 (500 + 2)
	ReduceInt(as, div) // Output: 250 (500 / 2)
}

func Itoa(nb int) string {
	var result string
	for nb > 0 {
		digit := nb % 10
		result = string('0'+digit) + result
		nb /= 10
	}
	return result
}
