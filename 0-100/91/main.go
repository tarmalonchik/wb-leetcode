package main

import (
	"strconv"
)

func numDecodings(s string) int {
	var mp = make(map[int]int)
	return decodingRecursive(mp, s, 0)
}

func decodingRecursive(mp map[int]int, s string, position int) (resp int) {
	if len(s) == 0 {
		return 1
	}
	if len(s) == 1 {
		if isValid(s) {
			return 1
		}
		return 0
	}

	singleValid := isValid(s[:1])
	singleSubNum := 0
	if singleValid {
		val, ok := mp[position+1]
		if ok {
			singleSubNum = val
		} else {
			singleSubNum = decodingRecursive(mp, s[1:], position+1)
			mp[position+1] = singleSubNum
		}
	}

	doubleValid := isValid(s[:2])
	doubleSubNum := 0
	if doubleValid {
		val, ok := mp[position+2]
		if ok {
			doubleSubNum = val
		} else {
			doubleSubNum = decodingRecursive(mp, s[2:], position+2)
			mp[position+2] = doubleSubNum
		}
	}

	if !singleValid && !doubleValid {
		return 0
	}

	if singleValid {
		resp += singleSubNum
	}
	if doubleValid {
		resp += doubleSubNum
	}

	return resp
}

func isValid(s string) bool {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid string")
	}
	if len(s) == 2 && val < 10 {
		return false
	}
	if val > 0 && val < 27 {
		return true
	}
	return false
}
