package main

import (
	"bytes"
)

func addSpaces(s string, spaces []int) string {
	outData := bytes.Buffer{}

	pos1 := 0
	pos2 := 0
	spacesIdx := 0

	for {
		if spacesIdx > len(spaces)-1 {
			outData.WriteString(s[pos1:])
			break
		}
		if pos2 == spaces[spacesIdx] {
			outData.WriteString(s[pos1:pos2])
			outData.WriteString(" ")
			pos1 = pos2
			spacesIdx++
		}
		pos2++
	}

	return outData.String()
}
