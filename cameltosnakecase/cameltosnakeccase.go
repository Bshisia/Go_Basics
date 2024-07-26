package main

import (
	"fmt"
)

// CamelToSnakeCase converts camelCase strings to snake_case without changing the case of letters
func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Check for invalid camelCase
	if !isValidCamelCase(s) {
		return s
	}

	var result []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i > 0 {
				result = append(result, '_') 
			}
			result = append(result, s[i])
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// isValidCamelCase checks if the string is valid camelCase
func isValidCamelCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	prevUpper := false
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			return false
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			if prevUpper || i == len(s)-1 {
				return false
			}
			prevUpper = true
		} else {
			prevUpper = false
		}
	}
	return true
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
