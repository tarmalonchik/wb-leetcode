package main

func minEnd(n int, x int) int64 {
	freeBits := getFreeBits(uint64(x))
	return int64(getNext(freeBits, uint64(x), uint64(n-1), getMaxBit(uint64(n-1))))
}

func getNext(freeBits []int, x, n uint64, maxSymbols uint8) (out uint64) {
	out = x
	for i := uint8(0); i < maxSymbols; i++ {
		nBit := getBit(n, i)
		if !nBit {
			continue
		}
		out = out | (1 << freeBits[i])
	}
	return out
}

func getMaxBit(num uint64) uint8 {
	for i := uint8(0); i < 64; i++ {
		if setBit(0, i) > num {
			return i
		}
	}
	return 0
}

func setBit(num uint64, pos uint8) uint64 {
	return num | (1 << pos)
}

func getFreeBits(x uint64) []int {
	size := 64
	out := make([]int, 0)
	for i := 0; i < size; i++ {
		if !getBit(x, uint8(i)) {
			out = append(out, i)
		}
	}
	return out
}

func getBit(num uint64, pos uint8) bool {
	if num&(1<<pos) > 0 {
		return true
	}
	return false
}
