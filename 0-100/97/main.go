package main

import (
	"fmt"
)

func main() {
	//s1 := "abababababababababababababababababababababababababababababababababababababababababababababababababbb"
	//s2 := "babababababababababababababababababababababababababababababababababababababababababababababababaaaba"
	//s3 := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"

	s1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s3 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	fmt.Println(isInterleave(s1, s2, s3))
}

type chain struct {
	s1pos, s2pos, s3pos int
	nextItem            *chain
}

func addHeadToChain(c *chain, s1pos, s2pos, s3pos int) *chain {
	if c == nil {
		return &chain{
			s1pos: s1pos,
			s2pos: s2pos,
			s3pos: s3pos,
		}
	}

	newItem := &chain{
		s1pos: s1pos,
		s2pos: s2pos,
		s3pos: s3pos,
	}
	newItem.nextItem = c
	return newItem
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	var head *chain
	secondTrack := false

	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	if s3 == "" {
		return s1 == "" && s2 == ""
	}

	s1pos := 0
	s2pos := 0
	s3pos := 0

	for {
		if (s3pos > len(s3)-1) || (s1pos > len(s1)-1) || (s2pos > len(s2)-1) {
			if checkIfValid(s1, s2, s3, s1pos, s2pos, s3pos) {
				return true
			}

			if head != nil {
				secondTrack = true
				head, s1pos, s2pos, s3pos = getFromHead(head)
				continue
			}
			return false
		}

		if s1[s1pos] == s2[s2pos] {
			if s3[s3pos] != s1[s1pos] {
				if head == nil {
					return false
				}

				secondTrack = true

				head, s1pos, s2pos, s3pos = getFromHead(head)
				continue
			}
			if !secondTrack {
				head = addHeadToChain(head, s1pos, s2pos, s3pos)
			}
			s3pos++
			if !secondTrack {
				s1pos++
				continue
			}
			s2pos++
			secondTrack = false
			continue
		}

		if s1[s1pos] == s3[s3pos] {
			s1pos++
			s3pos++
			continue
		}
		if s2[s2pos] == s3[s3pos] {
			s2pos++
			s3pos++
			continue
		}

		if head != nil {
			secondTrack = true
			head, s1pos, s2pos, s3pos = getFromHead(head)
			continue
		}
		return false
	}
}

func getFromHead(head *chain) (newHead *chain, s1, s2, s3 int) {
	s1, s2, s3 = head.s1pos, head.s2pos, head.s3pos
	return head.nextItem, s1, s2, s3
}

func checkIfValid(s1, s2, s3 string, s1pos, s2pos, s3pos int) bool {
	if s3pos > len(s3)-1 {
		return s1pos > len(s1)-1 && s2pos > len(s2)-1
	}
	if s1pos > len(s1)-1 {
		return s3[s3pos:] == s2[s2pos:]
	}
	if s2pos > len(s2)-1 {
		return s3[s3pos:] == s1[s1pos:]
	}
	return false
}
