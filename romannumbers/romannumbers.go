package main

import (
	"fmt"
	"os"
	"strconv"
)

func romanNumbers(n int) (string, string) {
	romanStr := ""
	romanCalc := ""

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	Calc := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(I-V)", "I"}

	for i, val := range values {
		for n >= val {
			n -= val
			romanStr += romanNums[i]
			if romanCalc != "" {
				romanCalc += "+"
			}
			romanCalc += Calc[i]
		}
	}
	return romanStr, romanCalc
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, calculation := romanNumbers(num)
	fmt.Println(calculation)
	fmt.Println(roman)
}
