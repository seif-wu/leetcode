package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(gcd(1, 4))
	fmt.Println(simplifiedFractions(9))
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

func simplifiedFractions(n int) []string {
	var r []string
	if n == 1 {
		return r
	}

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				r = append(r, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}

	return r
}
