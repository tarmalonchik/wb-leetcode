package main

import (
	"bytes"
	"strconv"
)

func compressedString(word string) string {
	var buffer bytes.Buffer
	prevSymbol := byte('*')
	counter := 0
	for i := 0; i < len(word); i++ {
		if (prevSymbol != word[i] && counter > 0) || counter >= 9 {
			buffer.WriteString(strconv.Itoa(counter) + string(prevSymbol))
			counter = 1
		} else {
			counter++
		}
		prevSymbol = word[i]
	}
	if counter > 0 {
		buffer.WriteString(strconv.Itoa(counter) + string(prevSymbol))
	}

	return buffer.String()
}
