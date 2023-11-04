package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("abbbaabccbaabacab", "ab*b*b*bc*ac*.*bb*"))
}

// a b*b*b* b c* a c* .*           bb*
// a     bb b    a    abccbaabaca  b

//	ab*b*b*bc*ac*.*bb* b* abbbaabccbaabacab
//	 b*b*b*bc*ac*.*bb* b*  bbbaabccbaabacab
//	   b*b*bc*ac*.*bb* b*  bbbaabccbaabacab
//	     b*bc*ac*.*bb* b*  bbbaabccbaabacab
//	       bc*ac*.*bb* b*  bbbaabccbaabacab
//	        c*ac*.*bb* b*   bbaabccbaabacab
//	          ac*.*bb* b*   bbaabccbaabacab

func isMatch(s string, p string) (resp bool) {
	lToken, rToken := regularExprToTokens(p)
	return internalMatcher(s, lToken, rToken)
}

func internalMatcher(s string, lToken, rToken *token) (resp bool) {
	if lToken == rToken {
		return lToken.matchSqueeze(s)
	}

	if priorityIsLeft(lToken, rToken) {
		if lToken.single {
			s, lToken, rToken, resp = moveLeftSide(s, lToken, rToken)
			if lToken == nil && rToken == nil {
				return resp
			}
		} else {
			lToken = lToken.next
			lToken.stealPrevValue()
		}
	} else {
		if rToken.single {
			s, lToken, rToken, resp = moveRightSide(s, lToken, rToken)
			if lToken == nil && rToken == nil {
				return resp
			}
		} else {
			rToken = rToken.prev
			rToken.stealNextValue()
		}
	}
	return internalMatcher(s, lToken, rToken)
}

func moveLeftSide(s string, lToken, rToken *token) (string, *token, *token, bool) {
	lPos, rPos := 0, len(s)-1
	startPosition := lToken.prev.copy()

	for {
		if !lToken.single {
			return s[lPos : rPos+1], lToken, rToken, false
		}

		if lPos > rPos {
			return "", nil, nil, lToken.ballast(rToken)
		}

		if lPos >= 0 && lToken.snuffOutPrev(s[lPos]) {
			lPos++
			continue
		}

		if lToken == rToken {
			return "", nil, nil, lToken.matchSqueeze(s[lPos : rPos+1])
		}

		if lToken.equal(s[lPos]) {
			lPos++
			lToken = lToken.next
			continue
		}

		if startPosition != nil && startPosition.equal(s[lPos]) {
			lPos++
			lToken = startPosition.next
			continue
		}
		return "", nil, nil, false
	}
}

func moveRightSide(s string, lToken, rToken *token) (string, *token, *token, bool) {
	lPos, rPos := 0, len(s)-1
	endPosition := rToken.next.copy()

	for {
		if !rToken.single {
			return s[lPos : rPos+1], lToken, rToken, false
		}

		if lPos > rPos {
			return "", nil, nil, lToken.ballast(rToken)
		}

		if rPos < len(s) && rToken.snuffOutNext(s[rPos]) {
			rPos--
			continue
		}

		if lToken == rToken {
			return "", nil, nil, lToken.matchSqueeze(s[lPos : rPos+1])
		}

		if rToken.equal(s[rPos]) {
			rPos--
			rToken = rToken.prev
			continue
		}

		if endPosition != nil && endPosition.equal(s[rPos]) {
			rPos--
			rToken = endPosition.prev
			continue
		}

		return "", nil, nil, false
	}
}

func (t *token) stealNextValue() {
	if t != nil && t.next != nil && t.next.getSymbol() != nil && !t.single {
		if *t.next.getSymbol() == anySymbol {
			t.value = anySymbol
		}
	}
}

func (t *token) stealPrevValue() {
	if t != nil && t.prev != nil && t.prev.getSymbol() != nil && !t.single {
		if *t.prev.getSymbol() == anySymbol {
			t.value = anySymbol
		}
	}
}

func (t *token) stringValue() string {
	if t == nil {
		return "null"
	}
	if t.single {
		return string(t.value)
	}
	return string(t.value) + string(anyCount)
}

func (t *token) matchSqueeze(in string) bool {
	fmt.Println(in)
	var (
		prevToken, nextToken *token
		lPos                 int
		rPos                 = len(in)
		lStuck, rStuck       bool
	)

	if t.prev != nil && !t.prev.single {
		prevToken = t.prev.copy()
	}

	if t.next != nil && !t.next.single {
		nextToken = t.next.copy()
	}

	if prevToken == nil {
		lStuck = true
	}
	if nextToken == nil {
		rStuck = true
	}

	for {
		if rPos-lPos == 1 && t.single {
			return t.equal(in[lPos])
		}
		if rPos == lPos && !t.single {
			return true
		}
		if lStuck && rStuck {
			if t.single {
				return false
			}
			for i := lPos; i < rPos; i++ {
				if !t.equal(in[i]) {
					return false
				}
			}
			return true
		}

		if !lStuck && !t.equal(in[lPos]) {
			if prevToken.equal(in[lPos]) {
				lPos++
			} else {
				lStuck = true
			}
			continue
		}

		if !rStuck && !t.equal(in[rPos-1]) {
			if nextToken.equal(in[rPos-1]) {
				rPos--
			} else {
				rStuck = true
			}
			continue
		}

		if !lStuck {
			if t.single {
				if lPos+1 < len(in) {
					if !prevToken.equal(in[lPos+1]) {
						lStuck = true
						continue
					}
				}
			}

			if prevToken.equal(in[lPos]) {
				lPos++
			} else {
				lStuck = true
			}
			continue
		}

		if !rStuck {
			if t.single {
				if rPos-1 > 0 {
					if !nextToken.equal(in[rPos-2]) {
						rStuck = true
						continue
					}
				}
			}

			if nextToken.equal(in[rPos-1]) {
				rPos--
			} else {
				rStuck = true
			}
			continue
		}
	}
}

func regularExprToTokens(input string) (lToken, rToken *token) {
	currentToken := &token{}
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == anyCount {
			currentToken.prev = &token{
				single: false,
				value:  input[i-1],
				next:   currentToken,
			}
			i--
		} else {
			currentToken.prev = &token{
				single: true,
				value:  input[i],
				next:   currentToken,
			}
		}
		currentToken = currentToken.prev
		if currentToken.next.value == 0 {
			rToken = currentToken
		}
	}
	return currentToken, rToken
}

type token struct {
	single     bool
	value      uint8
	prev       *token
	next       *token
	snuffedOut int
}

func (t *token) ballast(rToken *token) bool {
	tok := t.copy()
	if tok == nil && rToken == nil {
		return false
	}
	counter := 0
	for {
		counter += tok.snuffedOut
		if tok.single && counter == 0 {
			return false
		}
		counter--
		if tok == rToken {
			return true
		}
		tok = tok.next
		if tok == nil {
			panic("first token and last token error")
		}
	}
}

func (t *token) equal(value uint8) bool {
	if t == nil {
		return false
	}
	if t.value == anySymbol {
		return true
	}
	if value == t.value {
		return true
	}
	if t.snuffedOut > 0 {
		t.snuffedOut--
		return true
	}
	return false
}

func (t *token) getSymbol() *uint8 {
	if t == nil {
		return nil
	}
	if t.single {
		return nil
	}
	return &t.value
}

func (t *token) printReg() (resp string) {
	tCopy := t.copy()
	for {
		if tCopy == nil || tCopy.value == 0 {
			return resp
		}
		symbol := ""
		if !tCopy.single {
			symbol = "*"
		}
		resp += string(tCopy.value) + symbol
		tCopy = tCopy.next
	}
}

func (t *token) copy() *token {
	if t == nil {
		return nil
	}
	return &token{
		single:     t.single,
		value:      t.value,
		prev:       t.prev,
		next:       t.next,
		snuffedOut: t.snuffedOut,
	}
}

func (t *token) snuffOutPrev(symbol uint8) bool {
	if t == nil || t.value != anySymbol || t.prev == nil || t.prev.single || !t.prev.equal(symbol) {
		return false
	}
	t.snuffedOut++
	return true
}

func (t *token) snuffOutNext(symbol uint8) bool {
	if t == nil || t.value != anySymbol || t.next == nil || t.next.single || !t.next.equal(symbol) {
		return false
	}
	t.snuffedOut++
	return true
}

func (t *token) getWeight() (weight int) {
	if !t.single {
		weight += 100
	}
	if t.value == anySymbol {
		weight += 10
	}
	return weight
}

func priorityIsLeft(lToken, rToken *token) bool {
	if lToken.getWeight() > rToken.getWeight() {
		return false
	}
	return true
}
