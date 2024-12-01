package main

func minimizedMaximum(n int, quantities []int) int {
	maxQuantity := quantities[0]
	for i := range quantities {
		if quantities[i] > maxQuantity {
			maxQuantity = quantities[i]
		}
	}
	minQuantity := 0

	if n == 1 {
		return quantities[0]
	}

	for {
		if maxQuantity-minQuantity == 0 || maxQuantity-minQuantity == 1 {
			if canDistribute(n, minQuantity, quantities) {
				return minQuantity
			}
			if canDistribute(n, maxQuantity, quantities) {
				return maxQuantity
			}
			return 0
		}

		center := minQuantity + (maxQuantity-minQuantity)/2
		if canDistribute(n, center, quantities) {
			maxQuantity = center
		} else {
			minQuantity = center
		}
	}
}

func canDistribute(stores, perStore int, quantities []int) bool {
	if len(quantities) == 0 {
		return true
	}
	if perStore == 0 {
		return false
	}

	for i := range quantities {
		stores -= quantities[i] / perStore
		remainCount := quantities[i] % perStore
		if remainCount != 0 {
			stores -= 1
		}
		if stores < 0 {
			return false
		}
	}
	return true
}
