package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	mp := make(mpManager, len(s1)*len(s2))

	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return recursive(mp, s1, s2, s3, 0, 0)
}

func recursive(mp mpManager, s1 string, s2 string, s3 string, pos1, pos2 int) bool {
	if pos1 == len(s1) && pos2 == len(s2) {
		return true
	}

	if pos1 == len(s1) {
		if s2[pos2] == s3[pos1+pos2] {
			val, ok := mp.get(pos1, pos2+1)
			if ok {
				return val
			}
			newVal := recursive(mp, s1, s2, s3, pos1, pos2+1)
			mp.set(pos1, pos2+1, newVal)
			return newVal
		}
		return false
	}

	if pos2 == len(s2) {
		if s1[pos1] == s3[pos1+pos2] {
			val, ok := mp.get(pos1+1, pos2)
			if ok {
				return val
			}
			newVal := recursive(mp, s1, s2, s3, pos1+1, pos2)
			mp.set(pos1+1, pos2, newVal)
			return newVal
		}
		return false
	}

	var a, b bool

	if s1[pos1] == s3[pos1+pos2] {
		val, ok := mp.get(pos1+1, pos2)
		if ok {
			a = val
		} else {
			newVal := recursive(mp, s1, s2, s3, pos1+1, pos2)
			mp.set(pos1+1, pos2, newVal)
			a = newVal
		}
	}

	if a == true {
		return a
	}

	if s2[pos2] == s3[pos1+pos2] {
		val, ok := mp.get(pos1, pos2+1)
		if ok {
			b = val
		} else {
			newVal := recursive(mp, s1, s2, s3, pos1, pos2+1)
			mp.set(pos1, pos2+1, newVal)
			b = newVal
		}
	}

	return b
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
