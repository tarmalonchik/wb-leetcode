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
			name:    "4",
			value:   "a",
			pattern: "b",
			match:   false,
		},
		{
			name:    "5",
			value:   "abba",
			pattern: "a.*a",
			match:   true,
		},
		{
			name:    "6",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*mn.*s",
			match:   true,
		},
		{
			name:    "7",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*m.*n.*s",
			match:   true,
		},
		{
			name:    "8",
			value:   "abbbbamnsdgdgas",
			pattern: "a.*mnn.*s",
			match:   false,
		},
		{
			name:    "9",
			value:   "uuuuu",
			pattern: "u*",
			match:   true,
		},
		{
			name:    "10",
			value:   "aab",
			pattern: "c*a*b",
			match:   true,
		},
		{
			name:    "11",
			value:   "aa",
			pattern: "aa",
			match:   true,
		},
		{
			name:    "12",
			value:   "aaa",
			pattern: "a.a",
			match:   true,
		},
		{
			name:    "13",
			value:   "aaba",
			pattern: "ab*a*c*a",
			match:   false,
		},
		{
			name:    "14",
			value:   "mississippi",
			pattern: "mis*is*ip*.",
			match:   true,
		},
		{
			name:    "15",
			value:   "mississippi",
			pattern: "mis*is*p*.",
			match:   false,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tCase.match, isMatch(tCase.value, tCase.pattern))
		})
	}
}
