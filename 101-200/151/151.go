package main

func reverseWords(s string) string {
	num := len(s)
	if num <= 0 {
		return ""
	}
	str := ""
	flag := true
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && flag {
			num = i
			continue
		}
		if s[i] == ' ' {
			str += s[i+1:num+1] + " "
			flag = true
		} else {
			if flag == true {
				num = i
			}
			flag = false
		}
	}
	if s[0] != ' ' {
		return str + s[:num+1]
	}
	if len(str) > 0 {
		return str[:len(str)-1]
	}
	return ""
}
