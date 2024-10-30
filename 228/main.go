package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	var resp []string

	t := token{}

	for i := range nums {
		out := t.add(nums[i])
		if out == "" {
			continue
		}
		resp = append(resp, out)
		t.clear()
		t.add(nums[i])
	}

	addVal := t.toString()

	if addVal != "" {
		resp = append(resp, addVal)
	}
	return resp
}

type token struct {
	wasSet    bool
	firstItem int
	lastItem  int
}

func (t *token) add(in int) string {
	if !t.wasSet {
		t.wasSet = true
		t.firstItem = in
		t.lastItem = in
		return ""
	} else {
		if t.lastItem+1 == in {
			t.lastItem = in
			return ""
		} else {
			return t.toString()
		}
	}
}

func (t *token) clear() {
	t.firstItem = 0
	t.lastItem = 0
	t.wasSet = false
}

func (t *token) toString() string {
	if !t.wasSet {
		return ""
	}
	if t.lastItem == t.firstItem {
		return strconv.Itoa(t.lastItem)
	}
	return strconv.Itoa(t.firstItem) + "->" + strconv.Itoa(t.lastItem)
}
