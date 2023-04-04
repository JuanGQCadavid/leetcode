package main

import "fmt"

func partitionString(s string) int {
	counter := 0
	subStr := make(map[rune]bool)

	for _, v := range s {
		if subStr[v] {
			counter++
			subStr = make(map[rune]bool)

		}
		subStr[v] = true
	}

	for _, v := range subStr {
		if v {
			counter++
			break
		}

	}
	return counter
}

func main() {
	input := "abacaba"
	fmt.Println("4 == ", partitionString(input))
}
