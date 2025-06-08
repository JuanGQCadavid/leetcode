package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	hashes := make(map[string][]int)

	for indx, word := range strs {
		sortedRune := []rune(word)
		slices.Sort(sortedRune)
		sortedWord := string(sortedRune)
		hashes[sortedWord] = append(hashes[sortedWord], indx)
	}

	result := make([][]string, 0, len(hashes))
	for _, indxs := range hashes {
		rest := make([]string, len(indxs))
		for idx, val := range indxs {
			rest[idx] = strs[val]
		}
		result = append(result, rest)
	}

	return result

}
