package main

import "fmt"

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	// Calculate the number of non-space characters
	nonSpaceCount := 0
	for _, char := range str {
		if char != ' ' {
			nonSpaceCount++
		}
	}

	if nonSpaceCount < 5 {
		return "Invalid Input\n"
	}

	var result []rune
	count := 0

	for i, char := range str {
		if char != ' ' {
			if count == 5 {
				count = 0
				result = append(result, ' ')
			}
			if (i+1)%6 == 0 { // Skip every 6th character
				continue
			}
			result = append(result, char)
			count++
		}
	}

	return string(result) + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
