package main

func decrypt(code []int, k int) []int {
	if len(code) == 0 {
		return nil
	}

	out := make([]int, 0, len(code))

	var pos1 = 0
	if k < 0 {
		k = abs(k)
		pos1 = len(code) - 1 - k
	}
	pos2 := pos1
	originalPos1 := pos1
	sum := 0

	for k > 0 {
		pos2 = next(len(code), pos2)
		sum += code[pos2]
		k--
		if k == 0 {
			break
		}
	}

	for {
		out = append(out, sum)

		pos2 = next(len(code), pos2)
		sum += code[pos2]
		pos1 = next(len(code), pos1)
		sum -= code[pos1]

		if pos1 == originalPos1 {
			break
		}

	}
	return out
}

func next(length, pos int) int {
	if pos == length-1 {
		return 0
	}
	return pos + 1
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
