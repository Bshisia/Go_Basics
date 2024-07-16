package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	var result [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)                      // []
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0) // new line
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3) // [[0 1 2] [3 4 5] [6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5) // [[0 1 2 3 4] [5 6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4) // [[0 1 2 3] [4 5 6 7]]
}
