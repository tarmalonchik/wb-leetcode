package main

func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}

	pos1 := 0
	pos2 := 0

	removalsPossible := len(str1) - len(str2)

	for {
		if pos2 == len(str2) {
			if pos1 == len(str1) {
				return true
			}
			if len(str1)-pos1 == removalsPossible {
				return true
			}
			return false
		}
		if str1[pos1] == str2[pos2] {
			pos1++
			pos2++
			continue
		}
		if getNextChar(str1[pos1]) == str2[pos2] {
			pos1++
			pos2++
			continue
		}

		if removalsPossible > 0 {
			removalsPossible--
			pos1++
			continue
		}
		return false
	}
}

func getNextChar(in byte) byte {
	if in == 'z' {
		return 'a'
	}
	return in + 1
}
