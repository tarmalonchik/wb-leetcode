package main

func myAtoi(str string) int {
	var result int32
	for i,_ := range str{
		if str[i] == 32 {
			continue
		}
		if str[i] == 45  {
			if len(str[i+1:]) > 0 && str[i+1] >=48 && str[i+1] <= 57 {
				result = -int32(str[i+1] - 48)
				if len(str[i+2:]) > 0{
					for _, value := range str[i+2:]{
						if value >= 48 && value <= 57{
							prev := result
							result *= 10
							if result/10 != prev{
								return -2147483648
							}
							prevv := result
							result -= int32(value - 48)
							if result > prevv{
								return -2147483648
							}
							continue
						}
						break
					}
					return int(result)
				}
				return int(result)
			}
		}
		if str[i] == 43  {
			if len(str[i+1:]) > 0 && str[i+1] >=48 && str[i+1] <= 57 {
				result = int32(str[i+1] - 48)
				if len(str[i+2:]) > 0{
					for _, value := range str[i+2:]{
						if value >= 48 && value <= 57{
							prev := result
							result *= 10
							if result/10 != prev{
								return 2147483647
							}
							prevv := result
							result += int32(value - 48)
							if result < prevv{
								return 2147483647
							}
							continue
						}
						break
					}
					return int(result)
				}
				return int(result)
			}
		}
		if str[i] >= 48 && str[i] <= 57{
			result = int32(str[i] - 48)
			if len(str[i+1:]) > 0 {
				for _, value := range str[i+1:]{
					if value >= 48 && value <= 57{
						prev := result
						result *= 10
						if result /10 != prev{
							return 2147483647
						}
						prevv := result
						result += int32(value - 48)
						if result < prevv{
							return 2147483647
						}
						continue
					}
					break
				}
				return int(result)
			}
			return int(result)
		}
		break
	}
	return 0
}

