package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	seenInStr2 := make(map[rune]bool)
	for _, char := range str2 {
		seenInStr2[char] = true
	}

	seenInStr1 := make(map[rune]bool)
	for _, char := range str1 {
		seenInStr1[char] = true
	}

	count := 0
	for char := range seenInStr1 {
		if !seenInStr2[char] {
			count++
		}
	}

	for char := range seenInStr2 {
		if !seenInStr1[char] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
