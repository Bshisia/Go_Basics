package main

import (
	"fmt"
	"os"
)

func Inter(str1, str2 string) string {
	seenInSec := make(map[rune]bool)
	for _, char := range str2 {
		seenInSec[char] = true
	}

	seenInRes := make(map[rune]bool)
	result := ""

	for _, char := range str1 {
		if seenInSec[char] && !seenInRes[char] {
			seenInRes[char] = true
			result += string(char)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	commonChars := Inter(str1, str2)
	fmt.Println(commonChars)
}
