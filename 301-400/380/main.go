package main

import (
	"fmt"
	"math/rand"
)

func main() {
	some := Constructor()
	some.Insert(0)
	some.Insert(1)
	fmt.Println(some.setData, some.positions)
	some.Remove(0)
	fmt.Println(some.setData, some.positions)
	some.Insert(2)
	some.Remove(1)
	fmt.Println(some.GetRandom())
}

type RandomizedSet struct {
	setData   map[int]int
	positions []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		setData: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.setData[val]
	if !ok {
		this.positions = append(this.positions, val)
		this.setData[val] = len(this.positions) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	valPosition, ok := this.setData[val]
	if !ok {
		return false
	}
	this.setData[this.positions[len(this.positions)-1]] = valPosition
	this.positions[valPosition] = this.positions[len(this.positions)-1]
	this.positions = this.positions[:len(this.positions)-1]
	delete(this.setData, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.positions[rand.Intn(len(this.positions))]
}
