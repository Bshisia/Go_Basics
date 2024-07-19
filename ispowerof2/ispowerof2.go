package main

import (
	"fmt"
	"os"
)

func ispowerOf2(num int) bool {
	if num <= 0 {
		return false
	}
	return (num & (num - 1)) == 0
}

func Atoi(str string) int {
	result := 0
	for _, char := range str {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result
}

func main() {
	args := os.Args[1]
	result := Atoi(args)
	if ispowerOf2(result) {
		fmt.Println("true")
		return
	} else {
		fmt.Println("False")
		return
	}
}
