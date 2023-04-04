package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	result := make([]int, len(spells))
	for i, v := range spells {

		l, m, r := 0, 0, len(potions)-1

		for l <= r {
			m = (l + r) / 2

			if int64(v*potions[m]) >= success {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		result[i] = len(potions) - l

	}
	return result
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := 7

	fmt.Printf("[4,0,3] == %+v \n", successfulPairs(spells, potions, int64(success)))

	spells = []int{3, 1, 2}
	potions = []int{8, 5, 8}
	success = 16

	fmt.Printf("[2,0,2] == %+v\n", successfulPairs(spells, potions, int64(success)))
}
