package main

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	profit := 0

	pos1 := 0
	pos2 := 1
	for {
		if prices[pos1] < prices[pos2] {
			profit = getMax(profit, prices[pos2]-prices[pos1])
		} else {
			pos1 = pos2
			pos2++

			if pos1 >= len(prices)-1 {
				return profit
			}
			continue
		}

		pos2++
		if pos2 == len(prices)-1 {
			continue
		}
		if pos2 == len(prices) {
			pos1++
			pos2 = pos1 + 1
		}
		if pos1 >= len(prices)-2 {
			return profit
		}
	}
}

func getMax(profit, newItem int) int {
	if newItem > profit {
		return newItem
	}
	return profit
}
