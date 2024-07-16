package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Rot13(str string) string {
	var result string
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+13)%26)
		} else if char >= 'A' && char <= 'Z' {
			result += string('A' + (char-'a'+13)%26)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	str := os.Args[1]
	result := Rot13(str)
	for _, res := range  result {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
