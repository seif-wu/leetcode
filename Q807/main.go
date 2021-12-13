package main

import "fmt"

func main() {
	fmt.Println(maxIncreaseKeepingSkyline(([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}})), 35)
	fmt.Println(maxIncreaseKeepingSkyline(([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})), 0)
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	length := len(grid)
	// 竖排最大值
	vertical := make([]int, 0)
	for i := 0; i < length; i++ {
		max := 0
		for j := 0; j < length; j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		vertical = append(vertical, max)
	}

	result := 0
	for _, v := range grid {
		max := 0
		for _, r := range v {
			if r > max {
				max = r
			}
		}

		for k, r := range v {
			if vertical[k] < max {
				result += vertical[k] - r
			} else {
				result += max - r
			}
		}
	}
	return result
}
