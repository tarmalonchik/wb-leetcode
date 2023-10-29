package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("abcbccbcbaabbcbb", "c*a.*ab*.*ab*a*..b*"))
}

func isMatch(s string, p string) (resp bool) {
	tokenOne, tokenTwo := regularExprToTokens(p)

	if tokenOne == tokenTwo {
		tokenOne.next = nil
		tokenOne.prev = nil
		return tokenOne.matchSqueeze(s)
	}

	posOne, posTwo := 0, len(s)-1

	if tokenOne.one {
		tokenOne, tokenTwo, resp = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
		if tokenOne == nil && tokenTwo == nil {
			return resp
		}
	}

	if tokenTwo.one {
		tokenOne, tokenTwo, resp = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
		if tokenOne == nil && tokenTwo == nil {
			return resp
		}
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func internalMatcher(s string, tokenOne, tokenTwo *token) (resp bool) {
	if tokenOne == tokenTwo {
		return tokenOne.matchSqueeze(s)
	}
	if tokenOne.next == tokenTwo {
		for i := range s {
			if tokenOne.equal(s[i]) {
				continue
			}
			if tokenTwo.equal(s[i]) {
				continue
			}
			return false
		}
		return true
	}

	tokenTwo = tokenTwo.prev
	tokenTwo.stealNextValue()

	tokenOne = tokenOne.next
	tokenOne.stealPrevValue()

	if tokenOne == tokenTwo {
		return tokenOne.matchSqueeze(s)
	}

	posOne, posTwo := 0, len(s)-1
	tokenOne, tokenTwo, resp = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
	if tokenOne == nil && tokenTwo == nil {
		return resp
	}
	tokenOne, tokenTwo, resp = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
	if tokenOne == nil && tokenTwo == nil {
		return resp
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func moveLeftSide(s string, firstToken, lastToken *token, posFirst, posLast *int) (*token, *token, bool) {
	startPosition := firstToken.prev.copy()

	if firstToken != nil && firstToken.prev != nil && firstToken.value == anySymbol && firstToken.prev.value == anySymbol {
		*posFirst++
		if *posFirst > *posLast {
			return nil, nil, firstToken.ballast(lastToken)
		}
		firstToken.one = false
		firstToken.value = firstToken.prev.value
	}

	for {
		if !firstToken.one {
			return firstToken, lastToken, false
		}

		if *posFirst > *posLast {
			return nil, nil, firstToken.ballast(lastToken)
		}

		if firstToken == lastToken {
			return nil, nil, firstToken.matchSqueeze(s[*posFirst : *posLast+1])
		}

		if firstToken.equal(s[*posFirst]) {
			*posFirst++
			firstToken = firstToken.next
			continue
		}

		if startPosition != nil && startPosition.equal(s[*posFirst]) {
			*posFirst++
			firstToken = startPosition.next
			continue
		}
		return nil, nil, false
	}
}

func moveRightSide(s string, firstToken, lastToken *token, posFirst, posLast *int) (*token, *token, bool) {
	endPosition := lastToken.next.copy()

	if lastToken != nil && lastToken.next != nil && lastToken.value == anySymbol && lastToken.next.value == anySymbol {
		*posLast--
		if *posFirst > *posLast {
			return nil, nil, firstToken.ballast(lastToken)
		}
		lastToken.one = false
		lastToken.value = lastToken.next.value
	}

	for {
		if !lastToken.one {
			return firstToken, lastToken, false
		}

		if *posFirst > *posLast {
			return nil, nil, firstToken.ballast(lastToken)
		}

		if firstToken == lastToken {
			return nil, nil, firstToken.matchSqueeze(s[*posFirst : *posLast+1])
		}

		if lastToken.equal(s[*posLast]) {
			*posLast--
			lastToken = lastToken.prev
			continue
		}

		if endPosition != nil && endPosition.equal(s[*posLast]) {
			*posLast--
			lastToken = endPosition.prev
			continue
		}

		return nil, nil, false
	}
}

func (t *token) stealNextValue() {
	if t.next.getSymbol() != nil && !t.one {
		if *t.next.getSymbol() == anySymbol {
			t.value = anySymbol
		}
	}
}

func (t *token) stealPrevValue() {
	if t.prev.getSymbol() != nil && !t.one {
		if *t.prev.getSymbol() == anySymbol {
			t.value = anySymbol
		}
	}
}

func (t *token) stringValue() string {
	if t.one {
		return string(t.value)
	}
	return string(t.value) + string(anyCount)
}

func (t *token) matchSqueeze(in string) bool {
	var (
		prevToken, nextToken  *token
		first                 int
		last                  = len(in)
		firstStuck, lastStuck bool
	)

	if t.prev != nil && !t.prev.one {
		prevToken = t.prev.copy()
	}

	if t.next != nil && !t.next.one {
		nextToken = t.next.copy()
	}

	if prevToken == nil {
		firstStuck = true
	}
	if nextToken == nil {
		lastStuck = true
	}

	for {
		if last-first == 1 && t.one {
			return t.equal(in[first])
		}
		if last == first && !t.one {
			return true
		}
		if firstStuck && lastStuck {
			if t.one {
				return false
			}
			for i := first; i < last; i++ {
				if !t.equal(in[i]) {
					return false
				}
			}
			return true
		}

		if !firstStuck && !t.equal(in[first]) {
			if prevToken.equal(in[first]) {
				first++
			} else {
				firstStuck = true
			}
			continue
		}

		if !lastStuck && !t.equal(in[last-1]) {
			if nextToken.equal(in[last-1]) {
				last--
			} else {
				lastStuck = true
			}
			continue
		}

		if !firstStuck {
			if t.one {
				if first+1 < len(in) {
					if !prevToken.equal(in[first+1]) {
						firstStuck = true
						continue
					}
				}
			}

			if prevToken.equal(in[first]) {
				first++
			} else {
				firstStuck = true
			}
			continue
		}

		if !lastStuck {
			if t.one {
				if last-1 > 0 {
					if !nextToken.equal(in[last-2]) {
						lastStuck = true
						continue
					}
				}
			}

			if nextToken.equal(in[last-1]) {
				last--
			} else {
				lastStuck = true
			}
			continue
		}
	}
}

func regularExprToTokens(input string) (firstToken, lastToken *token) {
	currentToken := &token{}
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == anyCount {
			currentToken.prev = &token{
				one:   false,
				value: input[i-1],
				next:  currentToken,
			}
			i--
		} else {
			currentToken.prev = &token{
				one:   true,
				value: input[i],
				next:  currentToken,
			}
		}
		currentToken = currentToken.prev
		if currentToken.next.value == 0 {
			lastToken = currentToken
		}
	}
	return currentToken, lastToken
}

type token struct {
	one   bool
	value uint8
	prev  *token
	next  *token
}

func (t *token) ballast(lastToken *token) bool {
	tok := t.copy()
	if tok == nil && lastToken == nil {
		return false
	}
	for {
		if tok.one {
			return false
		}
		if tok == lastToken {
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
	return value == t.value
}

func (t *token) getSymbol() *uint8 {
	if t == nil {
		return nil
	}
	if t.one {
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
		if !tCopy.one {
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
		one:   t.one,
		value: t.value,
		prev:  t.prev,
		next:  t.next,
	}
}
