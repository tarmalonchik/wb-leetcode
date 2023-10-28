package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("bbcacbabbcbaaccabc", "b*a*a*.c*bb*b*.*.*"))
}

func isMatch(s string, p string) bool {
	tokenOne, tokenTwo := regularExprToTokens(p)

	if tokenOne == tokenTwo {
		return matchSqueeze(tokenOne, s, nil, nil)
	}

	posOne, posTwo := 0, len(s)-1

	if tokenOne.one {
		var val *bool
		tokenOne, tokenTwo, val = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo, nil, nil)
		if val != nil {
			return *val
		}
	}

	if tokenTwo.one {
		var val *bool
		tokenOne, tokenTwo, val = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo, nil, nil)
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
		return matchSqueeze(tokenOne, s, &tokenOne.prev.value, &tokenOne.next.value)
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
		return matchSqueeze(tokenOne, s, &tokenOne.prev.value, &tokenOne.next.value)
	}

	posOne, posTwo := 0, len(s)-1
	var val *bool
	tokenOne, tokenTwo, val = moveLeftSide(s, tokenOne, tokenTwo, &posOne, &posTwo, &tokenOne.prev.value, &tokenTwo.next.value)
	if val != nil {
		return *val
	}

	tokenOne, tokenTwo, val = moveRightSide(s, tokenOne, tokenTwo, &posOne, &posTwo, &tokenOne.next.value, &tokenTwo.next.value)
	if val != nil {
		return *val
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func moveLeftSide(s string, firstToken, lastToken *token, posFirst, posLast *int, matchLeft, matchRight *uint8) (*token, *token, *bool) {
	var (
		matchLeftInternal  *uint8
		matchRightInternal *uint8
		rightWasApplied    bool
	)
	matchLeftInternal = copyPointerData(matchLeft)

	for {
		if !firstToken.one {
			if matchLeftInternal == nil {
				break
			}
			if firstToken.value == anySymbol {
				break
			}
			if firstToken.value == *matchLeftInternal {
				break
			}
			for {
				if *posFirst > *posLast {
					break
				}
				if equal(s[*posFirst], firstToken.value) {
					break
				}
				if equal(s[*posFirst], *matchLeftInternal) {
					*posFirst++
				} else {
					break
				}
			}
			break
		}

		if equal(s[*posFirst], firstToken.value) && !rightWasApplied {
			if matchLeftInternal != nil && !equal(s[*posFirst], *matchLeftInternal) {
				matchLeftInternal = nil
			}
			matchRightInternal = copyPointerData(matchRight)

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
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], nil, matchRightInternal)
				return nil, nil, &val
			}
		} else {
			if matchLeftInternal != nil {
				if equal(s[*posFirst], *matchLeftInternal) {
					*posFirst++
					continue
				}
			}
			if matchRightInternal != nil {
				rightWasApplied = true
				if equal(s[*posFirst], *matchRightInternal) {
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

func moveRightSide(s string, firstToken, lastToken *token, posFirst, posLast *int, matchLeft, matchRight *uint8) (*token, *token, *bool) {
	var (
		matchLeftInternal  *uint8
		matchRightInternal *uint8
		leftWasApplied     bool
	)
	matchRightInternal = copyPointerData(matchRight)

	for {
		if !lastToken.one {
			if matchRightInternal == nil {
				break
			}
			if lastToken.value == anySymbol {
				break
			}
			if lastToken.value == *matchRightInternal {
				break
			}
			for {
				if *posFirst > *posLast {
					break
				}
				if equal(s[*posLast], lastToken.value) {
					break
				}
				if equal(s[*posLast], *matchRightInternal) {
					*posLast--
				} else {
					break
				}
			}
			break
		}

		if equal(s[*posLast], lastToken.value) && !leftWasApplied {
			if matchRightInternal != nil && !equal(s[*posLast], *matchRightInternal) {
				matchRightInternal = nil
			}
			matchLeftInternal = copyPointerData(matchLeft)

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
				val = matchSqueeze(firstToken, s[*posFirst:*posLast+1], matchLeftInternal, nil)
				return nil, nil, &val
			}
		} else {
			if matchRightInternal != nil {
				if equal(s[*posLast], *matchRightInternal) {
					*posLast--
					continue
				}
			}

			if matchLeftInternal != nil {
				leftWasApplied = true
				if equal(s[*posLast], *matchLeftInternal) {
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

func copyPointerData(in *uint8) (out *uint8) {
	if in == nil {
		return out
	}
	outData := *in
	out = &outData
	return out
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
