package main

func isPalindrome(x int) bool {
	if x < 0{
		return false
	}
	bit := 1
	for{
		if x < bit*10{
			break
		}
		bit*=10
	}
	if bit == 1{
		return true
	}
	for ;bit != 0;{
		if x/bit == x%10{
			x = x - (bit+1)*(x%10)
			x /=10
			bit /= 100
		}else{
			return  false
		}
	}
	return true
}