package main

import "fmt"

func IsLetter(str string) bool {
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return true
}

func ZipString(s string) string {
	result := ""
	lens := len(s)
	if lens == 0 {
		return ""
	}

	i := 0
	for i < lens {
		// Skip non-alphabetic characters
		if !IsLetter(string(s[i])) {
			result += string(s[i])
			i++
			continue
		}

		// Count consecutive duplicate characters
		count := 1
		for j := i + 1; j < lens && s[j] == s[i]; j++ {
			count++
		}
		result += Itoa(count) + string(s[i])
		i += count // Move past the counted duplicates
	}
	return result
}

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	var result string
	for nb > 0 {
		digit := nb % 10
		result = string('0'+digit) + result
		nb /= 10
	}
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
