package main

import "fmt"

func PrintMemory(arr [10]byte) {
	//print the hexadecimal representstion of the bytes
	for i, b := range arr {
		fmt.Printf("%02X ", b)
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	//Print the ASCII representation of the bytes

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
