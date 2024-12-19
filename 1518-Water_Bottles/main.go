package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	var result int
	var emptyBottles int

	// first drink
	emptyBottles = numBottles
	result = numBottles

	for emptyBottles >= numExchange {
		// exchange
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		// drink
		emptyBottles += numBottles
		result += numBottles
	}

	return result
}

func main() {
	fmt.Println("Hello World")
}
