package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("caaacccbaababbb", "c*.*b*ba*ac*c*b*.*"))
}

func isMatch(s string, p string) bool {
	tokenOne, tokenTwo := regularExprToTokens(p)

	if tokenOne == tokenTwo {
		return matchSqueeze(tokenOne, s, nil, nil)
	}

	posOne, posTwo := 0, len(s)-1

	if tokenOne.one {
		var val *bool
		tokenOne, tokenTwo, val = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
		if val != nil {
			return *val
		}
	}

	if tokenTwo.one {
		var val *bool
		tokenOne, tokenTwo, val = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
		if val != nil {
			return *val
		}
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func internalMatcher(s string, tokenOne, tokenTwo *token) bool {
	if s == "" {
		return tokensAreBallast(tokenOne, tokenTwo)
	}

	if tokenOne == tokenTwo {
		return matchSqueeze(tokenOne, s, tokenOne.prev.getSymbol(), tokenOne.next.getSymbol())
	}
	if tokenOne.next == tokenTwo {
		var symbol = tokenOne.value
		for i := range s {
			if equal(s[i], symbol) {
				continue
			} else {
				symbol = tokenTwo.value
			}
			if equal(s[i], symbol) {
				continue
			}
			return false
		}
		return true
	}
	tokenOne = tokenOne.next
	tokenTwo = tokenTwo.prev
	tokenTwo = stealNextValue(tokenTwo)
	tokenOne = stealPrevValue(tokenOne)

	if tokenOne == tokenTwo {
		return matchSqueeze(tokenOne, s, tokenOne.prev.getSymbol(), tokenOne.next.getSymbol())
	}

	posOne, posTwo := 0, len(s)-1
	var val *bool
	tokenOne, tokenTwo, val = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
	if val != nil {
		return *val
	}

	tokenOne, tokenTwo, val = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo)
	if val != nil {
		return *val
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func moveLeftSide(s string, firstToken, lastToken *token, posFirst, posLast *int) (*token, *token, *bool) {
	startPosition := copyToken(firstToken.prev)
	allowRestart := true

	for {
		if *posFirst > *posLast {
			val := tokensAreBallast(firstToken, lastToken)
			return nil, nil, &val
		}

		if startPosition == nil || !equal(s[*posFirst], startPosition.value) {
			allowRestart = false
		}

		if !firstToken.one {
			break
		}
		if equal(s[*posFirst], firstToken.value) {
			firstToken = stealPrevMatcher(firstToken)

			firstToken = firstToken.next
			*posFirst++

			if *posFirst > *posLast {
				val := tokensAreBallast(firstToken, lastToken)
				return nil, nil, &val
			}

			if firstToken == lastToken {
				val := false
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], firstToken.prev.getSymbol(), firstToken.next.getSymbol())
				return nil, nil, &val
			}
		} else {
			if allowRestart {
				*posFirst++
				continue
			}
			val := false
			return nil, nil, &val
		}
	}
	return firstToken, lastToken, nil
}

func moveRightSide(s string, firstToken, lastToken *token, posFirst, posLast *int) (*token, *token, *bool) {
	endPosition := copyToken(lastToken.next)
	allowRestart := true

	for {
		if *posFirst > *posLast {
			val := tokensAreBallast(firstToken, lastToken)
			return nil, nil, &val
		}

		if endPosition == nil || !equal(s[*posLast], endPosition.value) {
			allowRestart = false
		}

		if !lastToken.one {
			break
		}
		if equal(s[*posLast], lastToken.value) {
			lastToken = stealNextMatcher(lastToken)

			lastToken = lastToken.prev
			*posLast--

			if *posFirst > *posLast {
				val := tokensAreBallast(firstToken, lastToken)
				return nil, nil, &val
			}

			if firstToken == lastToken {
				val := false
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], firstToken.prev.getSymbol(), firstToken.next.getSymbol())
				return nil, nil, &val
			}
		} else {
			if allowRestart {
				*posLast--
				continue
			}
			val := false
			return nil, nil, &val
		}
	}
	return firstToken, lastToken, nil
}

func stealNextValue(token *token) *token {
	if token.next.getSymbol() != nil && !token.one {
		if *token.next.getSymbol() == anySymbol {
			token.value = anySymbol
		}
	}
	return token
}

func stealPrevValue(token *token) *token {
	if token.prev.getSymbol() != nil && !token.one {
		if *token.prev.getSymbol() == anySymbol {
			token.value = anySymbol
		}
	}
	return token
}

func stealNextMatcher(token *token) *token {
	if token.next == nil {
		return token
	}
	if !token.next.one {
		if token.next.value == anySymbol {
			token.value = anySymbol
			token.one = false
		}
		if token.next.value == token.value {
			token.one = false
		}
	}
	return token
}

func stealPrevMatcher(token *token) *token {
	if token.prev == nil {
		return token
	}
	if !token.prev.one {
		if token.prev.value == anySymbol {
			token.value = anySymbol
			token.one = false
		}
		if token.prev.value == token.value {
			token.one = false
		}
	}
	return token
}

func tokensAreBallast(firstToken, lastToken *token) bool {
	if firstToken == nil && lastToken == nil {
		return false
	}
	for {
		if firstToken.one {
			return false
		}
		if firstToken == lastToken {
			return true
		}
		firstToken = firstToken.next
		if firstToken == nil {
			panic("first token and last token error")
		}
	}
}

func (t *token) stringValue() string {
	if t.one {
		return string(t.value)
	}
	return string(t.value) + string(anyCount)
}

func matchSqueeze2(token *token, in string, firstSqueeze, lastSqueeze *uint8) bool {
	for {

	}
}

func matchSqueeze(token *token, in string, firstSqueeze, lastSqueeze *uint8) bool {
	first := 0
	last := len(in)

	firstMatch := false
	lastMatch := false
	firstStuck := false
	lastStuck := false

	if firstSqueeze == nil {
		firstStuck = true
	}
	if lastSqueeze == nil {
		lastStuck = true
	}

	for {
		fmt.Println(first, last)
		if token.one {
			if first+1 == last {
				return equal(in[first], token.value)
			}
		} else {
			if first == last {
				return true
			}
		}

		if firstStuck && lastStuck {
			if token.one {
				return false
			}
			for i := first; i < last-1; i++ {
				if !equal(in[i], token.value) {
					return false
				}
			}
			return true
		}

		if !firstMatch && !firstStuck {
			firstMatch, firstStuck = moveSqueeze(token, in, &first, last, firstSqueeze, true)
			fmt.Println(firstMatch, firstStuck)
			continue
		}
		if !lastMatch && !lastStuck {
			fmt.Println("jangoggg")
			lastMatch, lastStuck = moveSqueeze(token, in, &last, first, lastSqueeze, false)
			fmt.Println(lastMatch, lastStuck)
			continue
		}

		if firstMatch && !firstStuck {
			first++
			firstMatch = false
			firstMatch, firstStuck = moveSqueeze(token, in, &first, last, firstSqueeze, true)
			continue
		}
		if lastMatch && !lastStuck {
			last--
			lastMatch = false
			lastMatch, lastStuck = moveSqueeze(token, in, &last, first, lastSqueeze, false)
			continue
		}
	}
}

func moveSqueeze(token *token, in string, position *int, posToCompare int, squeeze *uint8, left bool) (match bool, stuck bool) {
	if squeeze == nil {
		return equal(in[*position], token.value), true
	}
	for {
		fmt.Println(*position, posToCompare)
		fmt.Println(*position, posToCompare)
		fmt.Println(*position, posToCompare)
		if *position == posToCompare {
			return false, true
		}

		if equal(in[*position], token.value) {
			return true, !equal(in[*position], *squeeze)
		}

		if equal(in[*position], *squeeze) {
			if left {
				*position++
			} else {
				*position--
			}
		} else {
			return false, true
		}
	}
}

func equal(value, pattern uint8) bool {
	if pattern == anySymbol {
		return true
	}
	return value == pattern
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

func (t *token) getSymbol() *uint8 {
	if t == nil {
		return nil
	}
	if t.one {
		return nil
	}
	return &t.value
}

func printReg(tokenOne *token) (resp string) {
	var tokenCopy = tokenOne
	for {
		if tokenCopy == nil || tokenCopy.value == 0 {
			return resp
		}
		symbol := ""
		if !tokenCopy.one {
			symbol = "*"
		}
		resp += string(tokenCopy.value) + symbol
		tokenCopy = tokenCopy.next
	}
}

func copyToken(in *token) *token {
	if in == nil {
		return nil
	}
	return &token{
		one:   in.one,
		value: in.value,
		prev:  in.prev,
		next:  in.next,
	}
}
