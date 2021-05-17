package main

func convert(s string, numRows int) string {
	var (
		response   string
		arrPointer []int
		iterator   int
	)

	magicNumber := (numRows-2)*2 + 2

	if numRows == 1 || len(s) <= numRows {
		return s
	}

	for {
		arrPointer = append(arrPointer, iterator)
		arrPointer = append(arrPointer, iterator)
		if iterator < len(s) {
			response += string(s[iterator])
		}
		if iterator >= len(s) {
			break
		}
		iterator += magicNumber
	}

	for {
		for i := range arrPointer {
			if (i % 2) == 0 {
				arrPointer[i]--
			} else {
				arrPointer[i]++
			}
		}

		for i := range arrPointer {
			if i != 0 {
				if arrPointer[i] == arrPointer[i-1] {
					continue
				}
			}
			if arrPointer[i] < len(s) && arrPointer[i] >= 0 && (arrPointer[i] != arrPointer[i-1]) {
				response += string(s[arrPointer[i]])
			}
		}

		if arrPointer[1] == arrPointer[2] {
			break
		}
	}

	return response
}
