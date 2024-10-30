package main

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	} else if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	splitItems := split(abs(n))

	resp := float64(1)
	for i := range splitItems {
		var newNum = x
		for j := 0; j < splitItems[i]; j++ {
			newNum *= newNum
		}
		resp *= newNum
	}
	if n < 0 {
		resp = 1 / resp
	}
	return resp
}

func split(in int) (resp []int) {
	if in == 0 {
		return nil
	}
	for {
		add, residue := numTo2Power(in)
		resp = append(resp, add)
		if residue == 0 {
			break
		}
		in = residue
	}
	return resp
}

func numTo2Power(in int) (power int, residue int) {
	if in == 0 || in == 1 {
		return 0, 0
	}
	var sum = 1
	for {
		if sum > in {
			return power - 1, residue
		}
		residue = in - sum
		sum *= 2
		power++
	}
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
