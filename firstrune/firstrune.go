package main

import "github.com/01-edu/z01"

func FirstRune(str string) rune {
	h := []rune(str)
	return h[0]
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
