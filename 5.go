package main

func isPalindromeString(s string) bool{
	i := 0
	j := len(s)-1
	for ; i <= j ;{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	second := len(s)-1
	for ; second > 0; {
		first := 0
		item := second
		for ; item <= len(s) -1 ; {
			if isPalindromeString(s[first:item+1]){
				return s[first:item+1]
			}
			first ++
			item ++
		}
		second --
	}
	return s[0:second+1]
}
