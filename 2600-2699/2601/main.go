package main

func primeSubOperation(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	maxNum := 0
	for i := range nums {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	counter := newPrimeNumbersCounter(maxNum)
	prevNum := 1
	if nums[0] > 1 {
		closest := counter.getClosestPrime(nums[0] - 1)
		if closest == noPreviousPrime {
			prevNum = nums[0]
		} else {
			prevNum = nums[0] - int(closest)
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] <= prevNum {
			return false
		}
		offset := counter.getClosestPrime(nums[i] - (prevNum + 1))
		if offset == noPreviousPrime {
			prevNum = nums[i]
			continue
		}
		prevNum = nums[i] - int(offset)
	}
	return true
}

const noPreviousPrime = -2

type primeNumbersCounter struct {
	numbers []int16
}

func newPrimeNumbersCounter(maxNum int) *primeNumbersCounter {
	if maxNum < 2 {
		maxNum = 2
	}

	out := &primeNumbersCounter{}
	out.numbers = make([]int16, maxNum+1)
	out.numbers[0] = noPreviousPrime
	out.numbers[1] = noPreviousPrime

	currentIndex := int16(2)
	for {
		out.numbers[currentIndex] = currentIndex

		for i := currentIndex + currentIndex; i < int16(len(out.numbers)); i += currentIndex {
			out.numbers[i] = noPreviousPrime
		}

		cont := false
		for i := currentIndex + 1; i < int16(len(out.numbers)); i++ {
			if out.numbers[i] == 0 {
				out.numbers[i] = i
				currentIndex = i
				cont = true
				break
			}
		}
		if cont {
			continue
		}
		break
	}

	currentIndex = 2
	for i := 2; i < len(out.numbers); i++ {
		if out.numbers[i] == noPreviousPrime {
			out.numbers[i] = currentIndex
		} else {
			currentIndex = int16(i)
		}
	}
	return out
}

func (p *primeNumbersCounter) getClosestPrime(num int) int16 {
	if num < len(p.numbers) {
		val := p.numbers[num]
		if val == noPreviousPrime {
			return -2
		}
		return val
	}
	panic("too big num")
}
