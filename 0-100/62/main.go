package main

func uniquePaths(m int, n int) int {
	doNotRepeat := make(mpManager, m*n)

	if m > 0 && n > 0 {
		return uniquePathsRecursion(doNotRepeat, m, n) + 1
	}
	return 0
}

func uniquePathsRecursion(doNotRepeat mpManager, m int, n int) int {
	pathsCount := 0
	if m > 1 && n > 1 {
		pathsCount += 1

		mVal, mOK := doNotRepeat.get(m-1, n)
		if mOK {
			pathsCount += mVal
		} else {
			count := uniquePathsRecursion(doNotRepeat, m-1, n)
			doNotRepeat.set(m-1, n, count)
			pathsCount += count
		}

		nVal, nOK := doNotRepeat.get(m, n-1)
		if nOK {
			pathsCount += nVal
		} else {
			count := uniquePathsRecursion(doNotRepeat, m, n-1)
			doNotRepeat.set(m, n-1, count)
			pathsCount += count
		}
		return pathsCount
	}
	return pathsCount
}

type mpManager map[mpItem]int

func (m *mpManager) get(pos1, pos2 int) (value int, ok bool) {
	val, ok := (*m)[mpItem{
		pos1: uint16(pos1),
		pos2: uint16(pos2),
	}]
	return val, ok
}

func (m *mpManager) set(pos1, pos2 int, value int) {
	(*m)[mpItem{
		pos1: uint16(pos1),
		pos2: uint16(pos2),
	}] = value
}

type mpItem struct {
	pos1 uint16
	pos2 uint16
}
