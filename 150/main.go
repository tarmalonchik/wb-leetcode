package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stackItem := newStack(len(tokens))

	for i := len(tokens) - 1; i >= 0; i-- {
		currentItem := newItem(tokens[i])
		if currentItem.isOperator() {
			stackItem.add(currentItem)
			continue
		}
		stackLast, ok := stackItem.get(0)
		if ok && stackLast.isOperator() {
			stackItem.add(currentItem)
			continue
		}

		stackItem.add(currentItem)
		stackItem.tryToCompute()
	}

	for {
		if !stackItem.tryToCompute() {
			val, _ := stackItem.get(0)
			return val.value
		}
	}
}

const multiple = uint8(1)
const divide = uint8(2)
const sum = uint8(3)
const dec = uint8(4)

type item struct {
	operator uint8
	value    int
}

func (i item) apply(val, val2 int) int {
	switch i.operator {
	case multiple:
		return val * val2
	case divide:
		return val / val2
	case dec:
		return val - val2
	case sum:
		return val + val2
	}
	panic("invalid operator")
}

func (i item) isOperator() bool {
	return i.operator != 0
}

func newItem(in string) item {
	switch in {
	case "*":
		return item{operator: multiple}
	case "/":
		return item{operator: divide}
	case "-":
		return item{operator: dec}
	case "+":
		return item{operator: sum}
	}
	val, err := strconv.Atoi(in)
	if err != nil {
		panic("invalid number")
	}
	return item{value: val}
}

func newStack(maxCount int) stack {
	return stack{
		currentIndex: -1,
		items:        make([]item, maxCount),
	}
}

type stack struct {
	currentIndex int
	items        []item
}

func (s *stack) tryToCompute() bool {
	val, ok := s.get(0)
	if !ok {
		return false
	}
	if val.isOperator() {
		return false
	}

	val2, ok := s.get(1)
	if !ok {
		return false
	}
	if val2.isOperator() {
		return false
	}

	operator, ok := s.get(2)
	if !ok {
		return false
	}
	if !operator.isOperator() {
		return false
	}

	s.remove(2)
	s.add(item{value: operator.apply(val.value, val2.value)})
	s.tryToCompute()
	return true
}

func (s *stack) add(in item) {
	s.currentIndex++
	s.items[s.currentIndex] = in
}

func (s *stack) get(offset int) (val item, ok bool) {
	if s.currentIndex-offset < 0 {
		return item{}, false
	}
	return s.items[s.currentIndex-offset], true
}

func (s *stack) remove(offset int) bool {
	if s.currentIndex-offset < 0 {
		return false
	}
	s.currentIndex -= 1 + offset
	return true
}
