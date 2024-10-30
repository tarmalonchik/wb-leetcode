package main

func lengthOfLastWord(s string) int {
	var resp int
	var findSymbol bool

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			resp++
			findSymbol = true
			continue
		}
		if findSymbol {
			break
		}
	}
	return resp
}
