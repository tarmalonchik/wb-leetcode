package main

func rob(nums []int) int {
	s := storage{}
	for i := range nums {
		s.addNum(nums[i])
	}
	return s.total
}

type storage struct {
	total    int
	decIfAdd int
}

func (s *storage) addNum(num int) {
	if num == 0 {
		s.decIfAdd = 0
		return
	}
	if s.total == 0 {
		s.total += num
		s.decIfAdd = num
		return
	}
	if s.decIfAdd == 0 {
		s.total += num
		s.decIfAdd = num
		return
	}
	if s.total+num-s.decIfAdd > s.total {
		s.total = s.total + num - s.decIfAdd
		s.decIfAdd = num - s.decIfAdd
		return
	} else {
		s.decIfAdd = 0
	}
}
