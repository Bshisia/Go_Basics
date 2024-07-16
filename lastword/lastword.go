package main

import (
	"os"

	"github.com/01-edu/z01"
)

func LastWord(str string) string {
	var lastword string
	var start bool
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			start = true
			lastword = string(str[i]) + lastword
		} else if start {
			break
		}
	}
	return lastword
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	lastword := LastWord(str)
	if len(lastword) > 0 {
		for _, strs := range lastword {
			z01.PrintRune(strs)
		}
	}

}
