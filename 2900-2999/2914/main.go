package main

func minChanges(s string) int {
	var out int
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			out++
		}
	}
	return out
}
