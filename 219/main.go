package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]interface{}, k)
	if len(nums) < 2 {
		return false
	}

	pos1 := 0
	pos2 := 0

	for {
		if pos2 >= len(nums) {
			return false
		}
		if pos2-pos1 <= k {
			if _, ok := mp[nums[pos2]]; ok {
				return true
			}
			mp[nums[pos2]] = nil
			pos2++
		} else {
			delete(mp, nums[pos1])
			pos1++
		}
	}
}
