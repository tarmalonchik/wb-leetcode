package main

func addToMap(mp map[uint8]uint32, val uint8) {
	mp[val]++
}

func getFromMap(mp map[uint8]uint32, val uint8) bool {
	count, ok := mp[val]
	if !ok {
		return false
	}
	if count == 0 {
		return false
	}
	mp[val]--
	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[uint8]uint32, len(magazine))
	for i := range magazine {
		addToMap(mp, magazine[i])
	}

	for i := range ransomNote {
		if !getFromMap(mp, ransomNote[i]) {
			return false
		}
	}
	return true
}
