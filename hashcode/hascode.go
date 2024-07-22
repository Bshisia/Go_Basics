package main

import (
	"fmt"
)

// HashCode computes a hashed string based on the given formula.
func HashCode(dec string) string {
	length := len(dec)
	hashed := ""

	for i := 0; i < length; i++ {
		// Compute the hash value for the current character
		asciiValue := int(dec[i])
		hashValue := (asciiValue + length) % 127

		// Ensure the result is within the printable ASCII range
		if hashValue < 32 {
			hashValue += 33
		}

		// Append the character to the result
		hashed += string(hashValue)
	}

	return hashed
}

func main() {
	fmt.Println(HashCode("A"))          // Output: B
	fmt.Println(HashCode("AB"))         // Output: CD
	fmt.Println(HashCode("BAC"))        // Output: EDF
	fmt.Println(HashCode("Hello World")) // Output: Spwwz+bz}wo
}
