package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumDifference([]int{90}, 1))
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 3))
}

func minimumDifference(nums []int, k int) int {
	sort.IntSlice(nums).Sort()
	l := len(nums)
	if l <= 1 {
		return 0
	}
	r := nums[l-1]
	for index, v := range nums[:l-k+1] {
		tr := nums[index+k-1] - v
		if tr < r {
			r = tr
		}
	}
	return r
}
