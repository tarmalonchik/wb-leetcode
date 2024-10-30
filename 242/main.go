package main

func isAnagram(s string, t string) bool {
	sMp := make(map[uint8]uint32, len(s))
	tMp := make(map[uint8]uint32, len(s))

	for i := range s {
		sMp[s[i]]++
	}

	for i := range t {
		tMp[t[i]]++
	}

	if len(sMp) != len(tMp) {
		return false
	}

	for key, val := range sMp {
		num, ok := tMp[key]
		if !ok {
			return false
		}
		if val != num {
			return false
		}
	}
	return true
}
