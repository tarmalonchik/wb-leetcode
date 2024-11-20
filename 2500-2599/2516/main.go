package main

import (
	"math"
)

// sliding window approach
func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}

	mayBe := isPossible(s, k)
	if mayBe < 0 {
		return mayBe
	}
	if mayBe > 0 {
		return mayBe
	}

	multiplication := 0.5
	position := processRecursive(s, len(s), k, multiplication, -1)
	return len(s) - (position + 1)
}

func processRecursive(s string, pos2, k int, multiplication float64, maxPos int) int {
	data := abc{k, k, k}

	for i := pos2 + 1; i < len(s); i++ {
		data.dec(s[i])
	}

	pos1Copy := -1
	pos2Copy := pos2

	wasZero := false
	wasLess := false

	for {
		if data.isZero() {
			wasZero = true
		}

		if data.isLess() {
			wasLess = true
		}

		if pos2Copy >= len(s)-1 {
			break
		}

		pos1Copy++
		pos2Copy++

		data.dec(s[pos1Copy])
		data.inc(s[pos2Copy])
	}

	if wasLess {
		addition := int(math.Round(float64(len(s)) * multiplication))
		if addition == 0 {
			return pos2
		}

		currentPos := maxPos
		if pos2 > currentPos {
			currentPos = pos2
		}
		newPos := processRecursive(s, pos2+addition, k, multiplication/2, currentPos)

		return getMax(currentPos, newPos)
	}

	if wasZero {
		return pos2
	}

	addition := int(math.Round(float64(len(s)) * multiplication))
	if addition == 0 {
		if maxPos != -1 {
			return maxPos
		}
		return -1
	}
	return processRecursive(s, pos2-addition, k, multiplication/2, maxPos)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPossible(s string, k int) int {
	if len(s) == 0 {
		return -1
	}

	if len(s) <= k {
		return -1
	}

	data := abc{k, k, k}

	for i := 0; i < len(s); i++ {
		data.dec(s[i])
	}

	if data.isZero() {
		return len(s)
	}
	if data.isGross() {
		return -1
	}
	return 0
}

type abc struct {
	a, b, c int
}

func (a *abc) inc(input byte) {
	switch input {
	case 97:
		a.a++
	case 98:
		a.b++
	case 99:
		a.c++
	}
}

func (a *abc) dec(input byte) {
	switch input {
	case 97:
		a.a--
	case 98:
		a.b--
	case 99:
		a.c--
	}
}

func (a *abc) isGross() bool {
	if a.a > 0 || a.b > 0 || a.c > 0 {
		return true
	}
	return false
}

func (a *abc) isZero() bool {
	if a.a == 0 && a.c == 0 && a.b == 0 {
		return true
	}
	return false
}

func (a *abc) isLess() bool {
	if a.a <= 0 && a.b <= 0 && a.c <= 0 {
		return a.a+a.b+a.c < 0
	}
	return false
}

func (a *abc) copy() abc {
	return abc{
		a: a.a,
		b: a.b,
		c: a.c,
	}
}
