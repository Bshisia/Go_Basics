package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	if slice1 == nil && slice2 == nil {
		return []int(nil)
	}
	for _, Slc2 := range slice2 {
		slice1 = append(slice1, Slc2)
	}

	return slice1
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{3, 4, 5}, []int{6, 7, 8}))
	fmt.Println(ConcatSlice([]int{}, []int{1, 2, 3}))
	fmt.Println(ConcatSlice([]int{}, []int{}))
}
