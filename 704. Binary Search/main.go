package main

import (
	"fmt"
)

// import (
// 	//"math"
// 	"fmt"
// 	"time"
// )

func search(nums []int, target int) int {
	upperLimit := len(nums) - 1
	lowerLimit := 0

	for lowerLimit <= upperLimit {

		middlePoint := lowerLimit + ((upperLimit - lowerLimit) / 2)
		x := nums[middlePoint]

		switch {
		case x == target:
			return middlePoint
		case x < target:
			lowerLimit = middlePoint + 1
		default:
			upperLimit = middlePoint - 1
		}

		// if nums[middlePoint] == target {
		// 	return middlePoint
		// }

		// if nums[middlePoint] < target {
		// 	lowerLimit = middlePoint + 1

		// } else {
		// 	upperLimit = middlePoint - 1
		// }
	}

	return -1
}

// func search(nums []int, target int) int {

// 	left := 0
// 	right := len(nums) - 1

// 	for left <= right {

// 		pivot := left + (right-left)/2

// 		if nums[pivot] > target {
// 			right = pivot - 1
// 		} else if nums[pivot] < target {
// 			left = pivot + 1
// 		} else {
// 			return pivot
// 		}

// 	}

// 	return -1

// }

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	// target := 9
	// response := search(nums, target)
	// fmt.Println(response)

	for i, v := range nums {
		fmt.Println(i, v)
		response := search(nums, v)
		fmt.Println(response == i)
	}

	// fmt.Println(5 / 2)
}
