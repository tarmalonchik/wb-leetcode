package main

var phone = map[uint8][]string{
	'1': nil,
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
	'0': nil,
}

func letterCombinations(digits string) []string {
	c := combiner{}
	c.process(digits, "")
	return c.data
}

type combiner struct {
	data []string
}

func (c *combiner) process(digits, payload string) {
	if digits == "" {
		if payload != "" {
			c.data = append(c.data, payload)
		}
		return
	}
	for i := range phone[digits[0]] {
		newPayload := payload
		c.process(digits[1:], newPayload+phone[digits[0]][i])
	}
}
