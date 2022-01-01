package main

import "fmt"

func main() {
	original, m, n := []int{1, 2, 3, 4}, 2, 2
	fmt.Println(construct2DArray(original, m, n), "[[1 2] [3 4]]")
	fmt.Println(construct2DArray([]int{3}, 1, 2), "[]")
	fmt.Println(construct2DArray([]int{1, 2}, 1, 1), "[]")
}

func construct2DArray(original []int, m int, n int) [][]int {
	oLen := len(original)
	if oLen != m*n {
		return make([][]int, 0)
	}

	result := make([][]int, 0, m)
	t := 0
	for t < oLen {
		nextN := t + n
		result = append(result, original[t:nextN])
		t = nextN
	}

	return result
}
