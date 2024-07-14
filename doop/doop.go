package main

import (
	"os"
)

func IsOperend(str string) bool {
	Ops := []string{"-", "/", "+", "*", "%"}
	for _, operation := range Ops {
		if str == operation {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 4 {
		return
	} else {
		if IsOperend(os.Args[2]) == false {
			return
		} else {
			val1 := os.Args[1]
			oper := os.Args[2]
			val2 := os.Args[3]
			intVal1 := Atoi(val1)
			intVal2 := Atoi(val2)
			if oper == "+" {
				result := Itoa(intVal1 + intVal2)
				os.Stdout.WriteString(result)
			}
		}
	}
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

func Itoa(nb int) string {
	var result string
	for nb > 0 {
		digit := nb % 10
		result = string('0' + digit) + result
		nb /= 10
	}
	return result
}
