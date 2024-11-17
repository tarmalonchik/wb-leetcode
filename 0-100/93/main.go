package main

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	return restoreRecursive(s, 4)
}

func restoreRecursive(s string, number int) (out []string) {
	if number == 0 {
		return nil
	}
	if len(s) == 0 {
		return nil
	}

	maxNum, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}

	if number == 4 {
		if len(s) < 4 {
			return nil
		}
		if maxNum > 255255255255 {
			return nil
		}
	}
	if number == 3 {
		if len(s) < 3 {
			return nil
		}
		if maxNum > 255255255 {
			return nil
		}
	}
	if number == 2 {
		if len(s) < 2 {
			return nil
		}
		if maxNum > 255255 {
			return nil
		}
	}
	if number == 1 {
		if maxNum > 255 {
			return nil
		}
	}

	if number == 1 {
		if s[0] == '0' && len(s) > 1 {
			return nil
		}
		return []string{s}
	}

	oneDig := string(s[0])
	nestedData := restoreRecursive(s[1:], number-1)
	if nestedData != nil {
		for i := range nestedData {
			out = append(out, oneDig+"."+nestedData[i])
		}
	}
	if s[0] == '0' {
		return out
	}

	if len(s) <= 2 {
		return out
	}

	twoDig := s[:2]
	nestedData = restoreRecursive(s[2:], number-1)

	if nestedData != nil {
		for i := range nestedData {
			out = append(out, twoDig+"."+nestedData[i])
		}
	}

	if len(s) <= 3 {
		return out
	}

	threeDig := s[:3]
	threeDigInt, err := strconv.Atoi(threeDig)
	if err != nil {
		return out
	}
	if threeDigInt > 255 {
		return out
	}
	nestedData = restoreRecursive(s[3:], number-1)
	if nestedData != nil {
		for i := range nestedData {
			out = append(out, threeDig+"."+nestedData[i])
		}
	}

	return out
}
