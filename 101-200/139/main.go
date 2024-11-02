package main

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]interface{}, len(wordDict))
	for i := range wordDict {
		mp[wordDict[i]] = nil
	}

	doNotRepeat := make(mpManager, len(s)*len(s))

	return recursive(doNotRepeat, mp, s, 0, 1)
}

func recursive(doNotRepeat mpManager, mp map[string]interface{}, s string, pos1, pos2 int) bool {
	if len(s) == 0 {
		return true
	}

	for {
		_, ok := mp[s[pos1:pos2]]
		if ok {
			if pos2 == len(s) {
				return true
			}

			case1, case1OK := doNotRepeat.get(pos2, pos2+1)
			if !case1OK {
				case1 = recursive(doNotRepeat, mp, s, pos2, pos2+1)
				doNotRepeat.set(pos2, pos2+1, case1)
			}

			case2, case2OK := doNotRepeat.get(pos1, pos2+1)
			if !case2OK {
				case2 = recursive(doNotRepeat, mp, s, pos1, pos2+1)
				doNotRepeat.set(pos2, pos2+1, case2)
			}
			return case1 || case2
		} else {
			if pos2 == len(s) {
				return false
			}
			pos2++
		}
	}
}

type mpItem struct {
	pos1 uint16
	pos2 uint16
}

type mpManager map[mpItem]bool

func (m *mpManager) get(pos1, pos2 int) (value bool, ok bool) {
	val, ok := (*m)[mpItem{
		pos1: uint16(pos1),
		pos2: uint16(pos2),
	}]
	if ok {
		return val, true
	}
	return false, false
}

func (m *mpManager) set(pos1, pos2 int, value bool) {
	(*m)[mpItem{
		pos1: uint16(pos1),
		pos2: uint16(pos2),
	}] = value
}
