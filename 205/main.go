package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[uint8]uint8, len(t))
	reverseMap := make(map[uint8]uint8, len(t))

	for i := range s {
		if revMapVal, ok := reverseMap[t[i]]; ok {
			if revMapVal != s[i] {
				return false
			}
		}
		mapVal, ok := mp[s[i]]
		if ok {
			if mapVal == t[i] {
				continue
			}
			return false
		}
		mp[s[i]] = t[i]
		reverseMap[t[i]] = s[i]
	}
	return true
}
