package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}
	start := nbrs[0]
	end := len(a)

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	//handle the negatives
	if start < 0 {
		start = len(a) + start
	}

	if end < 0 {
		end = len(a) + end
	}

	//ensurin gindices are within bounds of the slice
	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}
	if start >= len(a) || end <= start {
		return nil
	}
	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))      // Expected: []string{"algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 2, 4))   // Expected: []string{"ascii", "package"}
	fmt.Printf("%#v\n", Slice(a, -3))     // Expected: []string{"ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, -2, -1)) // Expected: []string{"package"}
	fmt.Printf("%#v\n", Slice(a, 2, 0))   // Expected: []string(nil)
}
