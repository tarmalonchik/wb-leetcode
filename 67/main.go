package main

const (
	zero = uint8('0')
	one  = uint8('1')
)

func addBinary(a string, b string) string {
	out := make([]uint8, 0, getMaxLen(a, b)+1)
	addition := zero

	aPos := position(len(a) - 1)
	bPos := position(len(b) - 1)

	for {
		if aPos == -1 && bPos == -1 {
			if addition != zero {
				out = append(out, addition)
			}
			break
		}
		value, newAddition := plusSymbols(aPos.decAndGetSymbol(a), bPos.decAndGetSymbol(b))
		value, addition = plusSymbols(value, addition)
		addition = getMax(newAddition, addition)
		out = append(out, value)
	}
	reverse(out)
	return string(out)
}

type position int

func (p *position) decAndGetSymbol(in string) uint8 {
	if *p == -1 {
		return zero
	}
	*p--
	return in[*p+1]
}

func plusSymbols(a, b uint8) (value uint8, addition uint8) {
	if a == one && b == one {
		return zero, one
	}
	if a == one || b == one {
		return one, zero
	}
	return zero, zero
}

func getMax(a, b uint8) uint8 {
	if a == one || b == one {
		return one
	}
	return zero
}

func getMaxLen(a, b string) int {
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func reverse(in []uint8) {
	lPos := 0
	rPos := len(in) - 1
	for {
		if lPos >= rPos {
			break
		}
		lPosVal := in[lPos]
		in[lPos] = in[rPos]
		in[rPos] = lPosVal
		rPos--
		lPos++
	}
}
