package main

func maxCount(banned []int, n int, maxSum int) int {
	mp := make(map[int]interface{}, len(banned))

	for i := range banned {
		mp[banned[i]] = nil
	}

	currentSum := 0
	count := 0
	for i := 1; i <= n; i++ {
		if _, ok := mp[i]; ok {
			continue
		}
		if currentSum+i <= maxSum {
			count++
			currentSum += i
		} else {
			break
		}
	}
	return count
}
