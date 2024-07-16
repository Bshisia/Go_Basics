package main

import "os"

func WdMatch(str1, str2 string) bool {
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	return i == len(str1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	if WdMatch(str1, str2) {
		os.Stdout.WriteString(str1 + "\n")
	}
}
