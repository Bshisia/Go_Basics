package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Alphamirror(str string) string {
	result := ""
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			result += string('Z' - (char - 'A'))
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	str := os.Args[1]
	result := Alphamirror(str)
	for _, res := range result {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
