package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := Atoi(os.Args[1])
	TabMult(n)
}

func TabMult(n int) {
	for i := 1; i <= 9; i++ {
		line := Itoa(i) + " "+ "X" + " " + Itoa(n) + " " + "=" + " " + Itoa(i*n) + "\n"
		os.Stdout.WriteString(line)
	}
}

func Itoa(nb int) string {
	result := ""
	for nb > 0 {
		digit := nb % 10
		result = string('0'+digit) + result
		nb /= 10
	}
	return result
}

func Atoi(n string) int {
	result := 0
	for _, char := range n {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result
}
