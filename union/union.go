package main

import (
	"fmt"
	"os"
)

func Union(str1, str2 string) string {
	seen := make(map[rune]bool)
	result := ""

	for _, char1 := range str1 {
		if !seen[char1] {
			seen[char1] = true
			result += string(char1)
		}
	}
	for _, char2 := range str2 {
		if !seen[char2] {
			seen[char2] = true
			result += string(char2)
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

	uniqueChars := Union(str1, str2)
	fmt.Println(uniqueChars)
}
