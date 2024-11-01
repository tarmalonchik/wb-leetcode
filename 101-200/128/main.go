package main

func longestConsecutive(nums []int) int {
	var longest int

	mp := make(map[int]interface{}, len(nums))

	for i := range nums {
		mp[nums[i]] = nil
	}

	for {
		val, ok := getRand(mp)
		if !ok {
			return longest
		}

		pos1 := val
		pos2 := val

		for {
			if pos2-pos1+1 > longest {
				longest = pos2 - pos1 + 1
			}

			if _, ok = mp[pos2+1]; ok {
				delete(mp, pos2)
				pos2++
				continue
			}

			if _, ok = mp[pos1-1]; ok {
				delete(mp, pos1)
				pos1--
				continue
			}

			delete(mp, pos1)
			delete(mp, pos2)
			break
		}
	}
}

func getRand(mp map[int]interface{}) (out int, ok bool) {
	for key, _ := range mp {
		out = key
		break
	}
	if _, ok = mp[out]; ok {
		return out, true
	}
	return out, false
}
