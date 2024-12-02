package main

func isPrefixOfWord(sentence string, searchWord string) int {
	counter := 1

	i := 0
	for {
		if i == len(sentence) {
			break
		}
		if sentence[i] == ' ' {
			counter++
			i++
			continue
		}

		if i+len(searchWord) > len(sentence) {
			return -1
		}

		if searchWord == sentence[i:i+len(searchWord)] {
			return counter
		}

		j := i + 1
		for {
			if j >= len(sentence) {
				return -1
			}
			if sentence[j] == ' ' {
				counter++
				i = j + 1
				break
			}
			j++
		}
	}
	return -1
}
