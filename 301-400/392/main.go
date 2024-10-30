package main

func isSubsequence(s string, t string) bool {
	origPos := 0
	heapPos := 0

	for {
		if origPos > len(s)-1 {
			return true
		}
		if heapPos > len(t)-1 {
			return false
		}

		if s[origPos] == t[heapPos] {
			origPos++
			heapPos++
		} else {
			heapPos++
		}
	}
}
