package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	var (
		counterHash = make(map[int]int)
	)

	for _, num := range nums {
		counterHash[num] += 1
	}

	var (
		keys = make([]int, 0, len(counterHash))
	)

	for val, _ := range counterHash {
		keys = append(keys, val)
	}

	sort.Slice(keys, func(i, j int) bool {
		return counterHash[keys[i]] > counterHash[keys[j]]
	})

	return keys[0:k]

}
