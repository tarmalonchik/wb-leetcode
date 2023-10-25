package main

import (
	"fmt"
)

const (
	anySymbol = '.'
	anyCount  = '*'
)

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

// ssissipp s*is*p*
// ssissipp is*

func isMatch(s string, p string) bool {
	tokenOne, tokenTwo := regularExprToTokens(p)
	if tokenOne == tokenTwo {
		return matchTokenWithString(tokenOne, s)
	}

	posOne, posTwo := 0, len(s)-1

	if tokenOne.single {
		fmt.Println("jajaj")
		var val *bool
		tokenOne, tokenTwo, val = moveTokenAndPos(s, tokenOne, tokenTwo, &posOne, &posTwo, false, false)
		if val != nil {
			return *val
		}
	}

	if tokenTwo.single {
		fmt.Println("jajaj")

		var val *bool
		tokenTwo, tokenOne, val = moveTokenAndPos(s, tokenTwo, tokenOne, &posTwo, &posOne, true, false)
		if val != nil {
			return *val
		}
	}

	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func internalMatcher(s string, tokenOne, tokenTwo *token) bool {
	fmt.Println(s, string(tokenOne.value), string(tokenTwo.value))
	if tokenOne == tokenTwo {
		return matchTokenWithString(tokenOne, s)
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
		return matchTokenWithString(tokenOne, s)
	}

	posOne, posTwo := 0, len(s)-1

	var val *bool
	tokenOne, tokenTwo, val = moveTokenAndPos(s, tokenOne, tokenTwo, &posOne, &posTwo, false, true)
	if val != nil {
		return *val
	}

	tokenTwo, tokenOne, val = moveTokenAndPos(s, tokenTwo, tokenOne, &posTwo, &posOne, true, true)
	if val != nil {
		return *val
	}
	return internalMatcher(s[posOne:posTwo+1], tokenOne, tokenTwo)
}

func moveTokenAndPos(s string, tokenMain, tokenToCompare *token, posMain, posToCompare *int, reverse, skipIfNonMatch bool) (tokenMainResp, tokenToCompareResp *token, resp *bool) {
	for {
		if equal(s[*posMain], tokenMain.value) {
			if tokenMain.single {
				if !reverse {
					tokenMain = tokenMain.next
				} else {
					tokenMain = tokenMain.prev
					tokenMain.next = nil
				}
			}

			if !reverse {
				*posMain++
			} else {
				*posMain--
			}

			if !reverse {
				if *posMain > *posToCompare {
					val := tokensAreBallast(tokenMain, tokenToCompare)
					return tokenMain, tokenToCompare, &val
				}
			} else {
				if *posMain < *posToCompare {
					val := tokensAreBallast(tokenToCompare, tokenMain)
					return tokenMain, tokenToCompare, &val
				}
			}

			if tokenMain == tokenToCompare {
				val := false
				if !reverse {
					if skipIfNonMatch {
						val = matchAnyStringPartWithToken(tokenMain, s[*posMain:*posToCompare])
					} else {
						val = matchTokenWithString(tokenMain, s[*posMain:*posToCompare+1])
					}
				} else {
					if skipIfNonMatch {
						val = matchAnyStringPartWithToken(tokenMain, s[*posToCompare:*posMain])
					} else {
						val = matchTokenWithString(tokenMain, s[*posToCompare:*posMain+1])
					}
				}
				return tokenMain, tokenToCompare, &val
			}

			if !tokenMain.single {
				break
			}
		} else {
			if skipIfNonMatch {
				if !reverse {
					*posMain++
				} else {
					*posMain--
				}
				continue
			}
			if !tokenMain.single {
				tokenMain = tokenMain.next
				continue
			}

			val := false
			return tokenMain, tokenToCompare, &val
		}
	}
	return tokenMain, tokenToCompare, nil
}

func tokensAreBallast(firstToken, lastToken *token) bool {
	if firstToken == nil && lastToken == nil {
		return false
	}
	for {
		if firstToken.single {
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

func regularExprToTokens(input string) (firstToken, lastToken *token) {
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
			lastToken = currentToken
			currentToken.next = nil
		}
	}
	return currentToken, lastToken
}

type token struct {
	single bool
	value  uint8
	prev   *token
	next   *token
}

func matchTokenWithString(token *token, s string) bool {
	if token.single == true {
		if len(s) > 1 {
			return false
		}
		if len(s) == 0 {
			return false
		}
		if s[0] == token.value {
			return true
		}
	}
	if token.value == anySymbol {
		return true
	}
	for i := range s {
		if s[i] != token.value {
			return false
		}
	}
	return true
}

func matchAnyStringPartWithToken(token *token, s string) bool {
	if !token.single {
		return true
	}
	if len(s) == 0 {
		return false
	}
	if token.value == anySymbol {
		return true
	}
	for i := range s {
		if s[i] == token.value {
			return true
		}
	}
	return false
}

func equal(value, pattern uint8) bool {
	if pattern == anySymbol {
		return true
	}
	return value == pattern
}
