package main

func myAtoi(str string) int {
	var result uint32
	for i, _ := range str {
		if str[i] == ' ' {
			continue
		}
		if str[i] == '+' || str[i] == '-' || (str[i] >= '0' && str[i] <= '9') {
			var sign int16
			if str[i] == '+' {
				sign = 10
			} else if str[i] == '-' {
				sign = 11
			}
			var bitMask uint32
			bitMask = 2147483648
			var flag bool
			for _, value := range str[i+int(sign/10):] {
				if value >= '0' && value <= '9' {
					result *= 10
					result += uint32(value - '0')
					if flag || ((result & bitMask) != 0) {
						if sign == 11 {
							return -2147483648
						} else {
							return 2147483647
						}
					}
					if !flag {
						if result > 429496729 {
							flag = true
						}
					}
				} else {
					break
				}
			}
			if sign == 11 {
				return -int(result)
			}
			return int(result)
		}
		return 0
	}
	return 0
}
