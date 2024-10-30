package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var myPrice = -1
	var profit int

	for i := 0; i < len(prices); i++ {
		if i == len(prices)-1 {
			if myPrice >= 0 {
				profit += prices[i] - myPrice
				myPrice = -1
			}
			return profit
		}

		if prices[i] <= prices[i+1] {
			if myPrice < 0 {
				myPrice = prices[i]
			}
		} else {
			if myPrice >= 0 {
				profit += prices[i] - myPrice
				myPrice = -1
			}
		}
	}
	return profit
}
