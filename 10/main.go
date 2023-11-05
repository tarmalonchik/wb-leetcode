package main

const (
	anySymbol = '.'
	anyCount  = '*'
)

func isMatch(s string, p string) (resp bool) {
	lToken := regularExprToTokens(p)
	lToken.simplifyTokens()
	return internalMatcher(s, lToken)
}

func internalMatcher(s string, lToken *token) bool {
	if lToken == nil {
		if s == "" {
			return true
		}
		return false
	}

	if s == "" {
		if lToken.single {
			return false
		}
		return internalMatcher(s, lToken.next)
	}

	if lToken.single {
		if !lToken.equal(s[0]) {
			return false
		}
		return internalMatcher(s[1:], lToken.next)
	}
	return internalMatcher(s, lToken.next) || (lToken.equal(s[0]) && (internalMatcher(s[1:], lToken.next) || internalMatcher(s[1:], lToken)))
}

func regularExprToTokens(input string) (lToken *token) {
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
			currentToken.next = nil
		}
	}
	return currentToken
}

type token struct {
	single bool
	value  uint8
	prev   *token
	next   *token
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
	return false
}

func (t *token) simplifyTokens() {
	if t == nil {
		return
	}

	tok := &token{
		single: t.single,
		value:  t.value,
		next:   t.next,
		prev:   t.prev,
	}

	for {
		if tok.next == nil {
			break
		}
		if tok.includes(tok.next) {
			tok.next = tok.next.next
			continue
		}
		if tok.next.includes(tok) {
			*tok = *tok.next
			continue
		}
		*tok = *tok.next
	}

	if tok.prev == nil {
		*t = *tok
	}
}

func (t *token) includes(compareToken *token) bool {
	if t.single || compareToken.single {
		return false
	}
	if t.value == anySymbol {
		return true
	}
	if t.value == compareToken.value {
		return true
	}
	return false
}
