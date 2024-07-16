package main

import "fmt"

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(str string) int {
	return len([]rune(str))
}
