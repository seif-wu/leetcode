package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))
}

func countKDifference(nums []int, k int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if int(math.Abs(float64(nums[i])-float64(nums[j]))) == k {
				r++
			}
		}
	}
	return r
}
