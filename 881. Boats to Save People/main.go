package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {

	sort.Ints(people)

	l, r, count := 0, len(people)-1, 0

	for l <= r {
		count++
		if people[r]+people[l] <= limit {
			l++
		}

		r--
	}

	return count
}

func main() {
	people := []int{1, 2}
	limit := 3

	fmt.Println("1 ==", numRescueBoats(people, limit))

	people = []int{3, 2, 2, 1}
	limit = 3

	fmt.Println("3 ==", numRescueBoats(people, limit))

	people = []int{3, 5, 3, 4}
	limit = 5

	fmt.Println("4 ==", numRescueBoats(people, limit))
}
