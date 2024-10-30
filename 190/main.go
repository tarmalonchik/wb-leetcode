package main

func reverseBits(num uint32) uint32 {
	out := uint32(0)
	for i := 0; i < 32; i++ {
		if getBit(num, uint8(31-i)) {
			out = setBit(out, uint8(i))
		}
	}
	return out
}

func setBit(num uint32, pos uint8) uint32 {
	return num | (1 << pos)
}

func getBit(num uint32, pos uint8) bool {
	if num&(1<<pos) > 0 {
		return true
	}
	return false
}
