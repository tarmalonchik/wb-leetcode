package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("acaabbaccbbacaabbbb", "a*.*b*.*a*aa*a*"))
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
	fmt.Println(printReg(tokenOne), tokenTwo.stringValue(), s)

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
	prevToken := firstToken.prev
	nextToken := lastToken.next

	var rightWasApplied bool

	for {
		if *posFirst > *posLast {
			val := tokensAreBallast(firstToken, lastToken)
			return nil, nil, &val
		}

		if !firstToken.one {
			if firstToken.prev.getSymbol() == nil {
				break
			}
			if firstToken.value == anySymbol {
				break
			}
			if firstToken.value == *firstToken.prev.getSymbol() {
				break
			}
			for {
				if *posFirst > *posLast {
					val := tokensAreBallast(firstToken, lastToken)
					return nil, nil, &val
				}
				if equal(s[*posFirst], firstToken.value) {
					break
				}
				if equal(s[*posFirst], *firstToken.prev.getSymbol()) {
					*posFirst++
				} else {
					break
				}
			}
			break
		}

		if equal(s[*posFirst], firstToken.value) && !rightWasApplied {
			if prevToken.getSymbol() != nil && !equal(s[*posFirst], *prevToken.getSymbol()) {
				prevToken = nil
			}

			if firstToken.one {
				firstToken = firstToken.next
			}

			*posFirst++

			if *posFirst > *posLast {
				val := tokensAreBallast(firstToken, lastToken)
				return nil, nil, &val
			}

			if firstToken == lastToken {
				val := false
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], nil, nextToken.getSymbol())
				return nil, nil, &val
			}
		} else {
			if prevToken.getSymbol() != nil {
				if equal(s[*posFirst], *prevToken.getSymbol()) {
					*posFirst++
					continue
				}
			}
			if nextToken.getSymbol() != nil {
				rightWasApplied = true
				if equal(s[*posFirst], *nextToken.getSymbol()) {
					*posFirst++
					continue
				}
			}
			val := false
			return nil, nil, &val
		}
	}
	return firstToken, lastToken, nil
}

func moveRightSide(s string, firstToken, lastToken *token, posFirst, posLast *int) (*token, *token, *bool) {
	prevToken := firstToken.prev
	nextToken := lastToken.next

	var (
		leftWasApplied bool
	)

	for {
		if *posFirst > *posLast {
			val := tokensAreBallast(firstToken, lastToken)
			return nil, nil, &val
		}

		if !lastToken.one {
			if lastToken.next.getSymbol() == nil {
				break
			}
			if lastToken.value == anySymbol {
				break
			}
			if lastToken.value == *lastToken.next.getSymbol() {
				break
			}
			for {
				if *posFirst > *posLast {
					val := tokensAreBallast(firstToken, lastToken)
					return nil, nil, &val
				}
				if equal(s[*posLast], lastToken.value) {
					break
				}
				if equal(s[*posLast], *lastToken.next.getSymbol()) {
					*posLast--
				} else {
					break
				}
			}
			break
		}

		if equal(s[*posLast], lastToken.value) && !leftWasApplied {
			if nextToken.getSymbol() != nil && !equal(s[*posLast], *nextToken.getSymbol()) {
				nextToken = nil
			}

			if lastToken.one {
				lastToken = lastToken.prev
			}

			*posLast--

			if *posFirst > *posLast {
				val := tokensAreBallast(firstToken, lastToken)
				return nil, nil, &val
			}

			if firstToken == lastToken {
				val := false
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], prevToken.getSymbol(), nil)
				return nil, nil, &val
			}
		} else {
			if nextToken.getSymbol() != nil {
				if equal(s[*posLast], *nextToken.getSymbol()) {
					*posLast--
					continue
				}
			}

			if prevToken.getSymbol() != nil {
				leftWasApplied = true
				if equal(s[*posLast], *prevToken.getSymbol()) {
					*posLast--
					continue
				}
			}

			val := false
			return nil, nil, &val
		}
	}
	return firstToken, lastToken, nil
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

func matchSqueeze(token *token, in string, leftSqueeze, rightSqueeze *uint8) bool {
	if token.value == '.' && !token.one {
		return true
	}
	if len(in) == 0 {
		if token.one {
			return false
		}
		return true
	}

	first := 0
	last := len(in) - 1

	match := false
	meet := false

	leftStuck := false
	rightStuck := false
	for {
		if first > last {
			meet = true
			break
		}

		if leftSqueeze == nil {
			leftStuck = true
		}
		if rightSqueeze == nil {
			rightStuck = true
		}

		if !leftStuck {
			if equal(in[first], *leftSqueeze) {
				if equal(in[first], token.value) {
					match = true
				}
				first++
				continue
			} else {
				leftStuck = true
			}
		}

		if !rightStuck {
			if equal(in[last], *rightSqueeze) {
				if equal(in[last], token.value) {
					match = true
				}
				last--
				continue
			} else {
				rightStuck = true
			}
		}

		if rightStuck && leftStuck {
			if token.one {
				if len(in[first:last]) == 0 {
					match = equal(in[first], token.value)
				}
				break
			}
			if !findSymbolInString(in[first:last+1], token.value) {
				break
			}
			match = true
			break
		}
	}
	if !meet && match {
		return true
	}
	if meet && (match || !token.one) {
		return true
	}
	return false
}

func findSymbolInString(in string, symbol uint8) bool {
	for i := range in {
		if !equal(in[i], symbol) {
			return false
		}
	}
	return true
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
