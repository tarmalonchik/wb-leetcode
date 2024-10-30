package main

func climbStairs(n int) int {
	sto := storage{}
	for i := 0; i < n; i++ {
		sto.calculate()
	}
	return sto.sum
}

type storage struct {
	finishWithOne int
	finishWithTwo int
	sum           int
}

func (s *storage) calculate() {
	if s.sum == 0 {
		s.finishWithOne++
		s.sum++
		return
	}
	s.sum = s.finishWithOne + s.finishWithTwo
	s.sum += s.finishWithOne
	s.finishWithOne = s.finishWithOne + s.finishWithTwo
	s.finishWithTwo = s.finishWithOne - s.finishWithTwo
}
