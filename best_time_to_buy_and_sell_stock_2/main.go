package main

import "fmt"

func maxProfitMinMax(prices []int) int {
	profit := 0
	localMin := prices[0]
	localMax := prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] > localMax {
			localMax = prices[i]
		} else {
			profit += localMax - localMin
			localMax = prices[i]
			localMin = prices[i]
		}
	}

	profit += localMax - localMin

	return profit
}

func maxProfitTrends(prices []int) int {
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfitMinMax([]int{7, 1, 5, 3, 6, 4}))
}
