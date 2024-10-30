package main

func longestCommonPrefix(strs []string) string {
	var (
		response string
	)

	if len(strs) == 0 {
		return response
	}

	for j := range strs[0] {
		for i := range strs {
			if j+1 > len(strs[i]) {
				return response
			}
			if i > 0 && strs[i][j] != strs[i-1][j] {
				return response
			}
		}
		response += string(strs[0][j])
	}
	return response
}
