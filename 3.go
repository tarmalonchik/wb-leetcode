package main

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int,0)
	k := 0
	for i:= 0; i < len(s); i ++ {
		if _, ok := mp[s[i]]; !ok{
			mp[s[i]] = i
		}else{
			if k < len(mp){
				k = len(mp)
			}
			i = mp[s[i]]
			mp = make(map[byte]int,0)
		}
	}
	if k < len(mp){
		k = len(mp)
	}
	return k
}
