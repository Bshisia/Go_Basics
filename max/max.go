package main

import "fmt"

func max(nb []int) int {
	max := 0
	for _, num := range nb {
		if num >= max {
			max = num
		}
	}
	return max
}

func main() {
	n := []int{123, 140, 82, 95}
	result := max(n)
	fmt.Println(result)
}
