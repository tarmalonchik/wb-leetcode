package main

func checkIfExist(arr []int) bool {
	mp := make(map[int]interface{})

	for i := range arr {
		_, ok := mp[arr[i]]
		if ok {
			return true
		}

		if arr[i]%2 == 0 {
			mp[arr[i]/2] = nil
		}
		mp[arr[i]*2] = nil
	}
	return false
}
