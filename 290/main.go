package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	stringArray := strings.Split(s, " ")
	if len(stringArray) != len(pattern) {
		return false
	}

	mp := make(map[uint8]string, len(stringArray))
	reverseMap := make(map[string]uint8, len(stringArray))

	for i := range pattern {
		reverseVal, ok := reverseMap[stringArray[i]]
		if ok {
			if reverseVal != pattern[i] {
				return false
			}
		}

		val, ok := mp[pattern[i]]
		if !ok {
			mp[pattern[i]] = stringArray[i]
			reverseMap[stringArray[i]] = pattern[i]
			continue
		}
		if val != stringArray[i] {
			return false
		}
	}
	return true
}
