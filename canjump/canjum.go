package main

import "fmt"

func CanJump(slice []uint) bool {
	i := 0

	if len(slice) == 1 {
		return true
	}

	if len(slice) == 0 {
		return false
	}
	for i < len(slice) {
		step := int(slice[i])
		if step == 0 {
			return false
		}
		i += step
		if i >= len(slice)-1 {
			return true
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // Expected output: true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // Expected output: false

	input3 := []uint{0}
	fmt.Println(CanJump(input3)) // Expected output: true
}
