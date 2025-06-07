package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		originalHash = make(map[rune]int)
	)

	for _, word := range s {
		originalHash[word] += 1
	}

	for _, word := range t {
		originalHash[word] += -1
	}

	for _, counter := range originalHash {
		if counter != 0 {
			return false
		}
	}

	return true
}
