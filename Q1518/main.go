package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(9, 3), 13, "\n ---------")

	fmt.Println(numWaterBottles(15, 4), 19, "\n ---------")
	fmt.Println(numWaterBottles(5, 5), 6, "\n ---------")
	fmt.Println(numWaterBottles(2, 3), 2, "\n ---------")
	fmt.Println(numWaterBottles(6, 5), 7, "\n ---------")
}

func numWaterBottles(numBottles int, numExchange int) int {
	if numExchange > numBottles {
		return numBottles
	}

	drinkNum, emptyNum := numBottles, numBottles

	for emptyNum >= numExchange {
		a := emptyNum / numExchange
		b := emptyNum % numExchange
		emptyNum = a + b
		drinkNum += a
	}

	return drinkNum
}
