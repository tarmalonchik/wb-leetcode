package main

import (
	"strings"
)

type romans struct {
	num   int
	roman string
}

var list = []romans{
	{
		num:   1000,
		roman: "M",
	},
	{
		num:   900,
		roman: "CM",
	},
	{
		num:   500,
		roman: "D",
	},
	{
		num:   400,
		roman: "CD",
	},
	{
		num:   100,
		roman: "C",
	},
	{
		num:   90,
		roman: "XC",
	},
	{
		num:   50,
		roman: "L",
	},
	{
		num:   40,
		roman: "XL",
	},
	{
		num:   10,
		roman: "X",
	},
	{
		num:   9,
		roman: "IX",
	},
	{
		num:   5,
		roman: "V",
	},
	{
		num:   4,
		roman: "IV",
	},
	{
		num:   1,
		roman: "I",
	},
}

func intToRoman(num int) string {
	var resp string
	for i := range list {
		val := num / list[i].num
		resp += strings.Repeat(list[i].roman, val)
		num = num % list[i].num
	}
	return resp
}
