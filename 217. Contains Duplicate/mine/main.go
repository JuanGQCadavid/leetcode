package main

func containsDuplicate(nums []int) bool {

	var (
		hash = make(map[int]bool)
	)
	for _, val := range nums {
		if _, ok := hash[val]; ok {
			return true
		}
		hash[val] = true
	}

	return false
}
