package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SearchReplace(inputStr, oldchar, newchar string) string {
	if len(oldchar) != 1 || len(newchar) != 1 {
		return inputStr
	}

	result := ""
	for _, str := range inputStr {
		if string(str) == oldchar {
			result += newchar
		} else {
			result += string(str)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	input := os.Args[1]
	oldchar := os.Args[2]
	newchar := os.Args[3]

	result := SearchReplace(input, oldchar, newchar)
	for _, res := range result {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
