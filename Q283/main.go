// LINK https://leetcode-cn.com/problems/move-zeroes/
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums, []int{1, 3, 12, 0, 0})

	nums2 := []int{0, 1, 0}
	moveZeroes(nums2)
	fmt.Println(nums2, []int{1, 0, 0})
}

// func moveZeroes(nums []int) {
// 	length := len(nums)
// 	for i := length - 1; i >= 0; i-- {
// 		if nums[i] != 0 {
// 			continue
// 		}

// 		for j := i; j < length-1; j++ {
// 			nums[j], nums[j+1] = nums[j+1], nums[j]
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
