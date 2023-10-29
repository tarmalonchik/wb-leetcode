package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	for _, tCase := range []struct {
		name    string
		value   string
		pattern string
		match   bool
	}{
		{
			name:    "1",
			value:   "s",
			pattern: ".",
			match:   true,
		},
		{
			name:    "2",
			value:   "s",
			pattern: ".*",
			match:   true,
		},
		{
			name:    "3",
			value:   "s",
			pattern: ".*.*.*",
			match:   true,
		},
		{
			name:    "4",
			value:   "a",
			pattern: "a",
			match:   true,
		},
		{
			name:    "5",
			value:   "a",
			pattern: "b",
			match:   false,
		},
		{
			name:    "6",
			value:   "abba",
			pattern: "a.*a",
			match:   true,
		},
		{
			name:    "7",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*mn.*s",
			match:   true,
		},
		{
			name:    "8",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*m.*n.*s",
			match:   true,
		},
		{
			name:    "9",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*mnn.*s",
			match:   false,
		},
		{
			name:    "10",
			value:   "uuuuu",
			pattern: "u*",
			match:   true,
		},
		{
			name:    "11",
			value:   "aab",
			pattern: "c*a*b",
			match:   true,
		},
		{
			name:    "12",
			value:   "aa",
			pattern: "aa",
			match:   true,
		},
		{
			name:    "13",
			value:   "aaa",
			pattern: "a.a",
			match:   true,
		},
		{
			name:    "14",
			value:   "mississippi",
			pattern: "mis*is*ip*.",
			match:   true,
		},
		{
			name:    "15",
			value:   "aaca",
			pattern: "ab*a*c*a",
			match:   true,
		},
		{
			name:    "16",
			value:   "aaaaaaaaaaaaab",
			pattern: "a*a*a*a*a*a*a*a*a*b",
			match:   true,
		},
		{
			name:    "17",
			value:   "aaaaaaaaaaaaab",
			pattern: "a*a*a*a*a*a*a*a*a*b",
			match:   true,
		},
		{
			name:    "18",
			value:   "aa",
			pattern: "a",
			match:   false,
		},
		{
			name:    "19",
			value:   "abbaaaabaabbcba",
			pattern: "a*.*ba.*c*..a*.a*.",
			match:   true,
		},
		{
			name:    "20",
			value:   "bbcacbabbcbaaccabc",
			pattern: "b*a*a*.c*bb*b*.*.*",
			match:   true,
		},
		{
			name:    "21",
			value:   "ccccacaaccaaaaabac",
			pattern: "..c*b*c*a*b*ba.*",
			match:   false,
		},
		{
			name:    "22",
			value:   "aaa",
			pattern: "ab*a",
			match:   false,
		},
		{
			name:    "23",
			value:   "aaa",
			pattern: "ab*ac*a",
			match:   true,
		},
		{
			name:    "24",
			value:   "acaabbaccbbacaabbbb",
			pattern: "a*.*b*.*a*aa*a*",
			match:   false,
		},
		{
			name:    "25",
			value:   "baabbbaccbccacacc",
			pattern: "c*..b*a*a.*a..*c",
			match:   true,
		},
		{
			name:    "26",
			value:   "abbcbaabcccaaaaab",
			pattern: "c*ab*a*a*a*b*b*",
			match:   false,
		},
		{
			name:    "27",
			value:   "cbbbaccbcacbcca",
			pattern: "b*.*b*a*.a*b*.a*",
			match:   true,
		},
		{
			name:    "28",
			value:   "cabcbabbacabbbba",
			pattern: "b*.*aa.*c*c*aa*b*",
			match:   false,
		},
		{
			name:    "29",
			value:   "abba",
			pattern: ".*aa.*",
			match:   false,
		},
		{
			name:    "30",
			value:   "abcbccbcbaabbcbb",
			pattern: "c*a.*ab*.*ab*a*..b*",
			match:   true,
		},
		{
			name:    "31",
			value:   "acbbcbcbcbaaacaac",
			pattern: "ac*.a*ac*.*ab*b*ac",
			match:   false,
		},
		{
			name:    "32",
			value:   "abbbaabccbaabacab",
			pattern: "ab*b*b*bc*ac*.*bb*",
			match:   true,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tCase.match, isMatch(tCase.value, tCase.pattern))
		})
	}
}

func TestMatchSqueeze(t *testing.T) {
	d := uint8('d')
	s := uint8('s')
	p := uint8('p')
	a := uint8('a')
	c := uint8('c')

	leftAny := uint8(anySymbol)
	rightAny := uint8(anySymbol)

	for _, tCase := range []struct {
		name                      string
		token                     *token
		input                     string
		leftSqueeze, rightSqueeze *uint8
		match                     bool
	}{
		{
			name: "1",
			token: &token{
				one:   true,
				value: uint8('s'),
			},
			input: "some",
			match: false,
		},
		{
			name: "2",
			token: &token{
				one:   false,
				value: uint8(anySymbol),
			},
			input: "some",
			match: true,
		},
		{
			name: "3",
			token: &token{
				one:   false,
				value: uint8('v'),
			},
			input: "vvvvv",
			match: true,
		},
		{
			name: "4",
			token: &token{
				one:   false,
				value: uint8('v'),
			},
			leftSqueeze: &d,
			input:       "dvvvvv",
			match:       true,
		},
		{
			name: "5",
			token: &token{
				one:   false,
				value: uint8('v'),
			},
			leftSqueeze:  &d,
			rightSqueeze: &s,
			input:        "dvsss",
			match:        true,
		},
		{
			name: "6",
			token: &token{
				one:   false,
				value: anySymbol,
			},
			input: "dvvvvvsss",
			match: true,
		},
		{
			name: "7",
			token: &token{
				one:   true,
				value: uint8('b'),
			},
			leftSqueeze:  &d,
			rightSqueeze: &s,
			input:        "b",
			match:        true,
		},
		{
			name: "8",
			token: &token{
				one:   true,
				value: uint8('m'),
			},
			leftSqueeze:  &leftAny,
			rightSqueeze: &rightAny,
			input:        "dfgdgsdfgsdfmdtatatat",
			match:        true,
		},
		{
			name: "9",
			token: &token{
				one:   false,
				value: uint8('m'),
			},
			leftSqueeze:  &leftAny,
			rightSqueeze: &rightAny,
			input:        "dfgdgsdfgsdfdtatatat",
			match:        true,
		},
		{
			name: "10",
			token: &token{
				one:   false,
				value: uint8('s'),
			},
			leftSqueeze:  nil,
			rightSqueeze: &p,
			input:        "ssipp",
			match:        false,
		},
		{
			name: "10",
			token: &token{
				one:   false,
				value: uint8('s'),
			},
			leftSqueeze:  nil,
			rightSqueeze: &p,
			input:        "ssipp",
			match:        false,
		},
		{
			name: "11",
			token: &token{
				one:   false,
				value: uint8('a'),
			},
			leftSqueeze:  &a,
			rightSqueeze: &a,
			input:        "cbaabcccaaaaa",
			match:        false,
		},
		{
			name: "12",
			token: &token{
				one:   true,
				value: uint8('a'),
			},
			input: "aa",
			match: false,
		},
		{
			name: "13",
			token: &token{
				one:   true,
				value: uint8('a'),
			},
			rightSqueeze: &rightAny,
			leftSqueeze:  &a,
			input:        "aababbb",
			match:        true,
		},
		{
			name: "14",
			token: &token{
				one:   false,
				value: uint8('.'),
			},
			rightSqueeze: &c,
			leftSqueeze:  &leftAny,
			input:        "bbacabbbb",
			match:        true,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tCase.match, matchSqueeze(tCase.token, tCase.input, tCase.leftSqueeze, tCase.rightSqueeze))
		})
	}
}
