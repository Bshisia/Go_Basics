package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	var longer, shorter []int 
	if len(slice1) >= len(slice2) {
		longer = slice1
		shorter = slice2
	} else {
		longer = slice2
		shorter = slice1
	}

	var result []int
	i := 0
	for i < len(shorter) {
		result = append(result, longer[i])
		result = append(result, shorter[i])
		i++
	}
	result = append(result, longer[1:]...)
	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
