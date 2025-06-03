package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {

	// Making the index array
	indx := make([]int, len(nums))
	for i := range nums {
		indx[i] = i
	}

	// sorting the index Desc
	sort.Slice(indx, func(i, j int) bool {
		return nums[indx[i]] > nums[indx[j]]
	})

	var (
		delta  = target
		result = make([]int, 0, 2)
	)

	for i := range indx {
		if nums[indx[i]] > delta {
			continue
		}

		delta -= nums[indx[i]]
		result = append(result, indx[i])
		if len(result) == 2 {
			break
		}
	}

	return result

	// for i := range indx {
	//     if nums[i] > target {
	//         continue
	//     }

	//     delta := target - nums[i]

	//     matchIndex := -1
	//     for j := range indx[i:] {
	//         if j == i {
	//             continue
	//         }
	//         sub := delta - nums[j]
	//         if sub == 0 {
	//             matchIndex = j
	//             break
	//         }

	//         // // If it is bigger than 0 then it meas that from now an on
	//         // // no number will make it equals to 0
	//         // if sub > 0 {
	//         //     break
	//         // }

	//     }

	//     if matchIndex != -1 {
	//         return []int{i, matchIndex}
	//     }
	// }
	return indx //[]int{-1,-1}
}
