package main

import (
	"strings"
)

type itemType uint8

const (
	dot      = itemType(2)
	skipper  = itemType(3)
	fileName = itemType(4)
)

type itemStorage struct {
	itemType  itemType
	itemValue string
}

func (i *itemStorage) Parse(in string) {
	in = strings.ReplaceAll(in, " ", "")
	switch in {
	case ".":
		i.itemType = dot
		i.itemValue = ""
	case "..":
		i.itemType = skipper
		i.itemValue = ""
	default:
		i.itemType = fileName
		i.itemValue = in
	}
}

func simplifyPath(path string) string {
	items := strings.Split(path, "/")
	itemsTypes := make([]itemStorage, len(items))
	for i := range items {
		itemsTypes[i].Parse(items[i])
	}

	out := make([]byte, len(path))

	skipCount := 0
	finishIdx := len(out)

	for i := len(itemsTypes) - 1; i >= 0; i-- {
		switch itemsTypes[i].itemType {
		case dot:
			continue
		case skipper:
			skipCount++
			continue
		case fileName:
			if itemsTypes[i].itemValue == "" {
				continue
			}

			if skipCount > 0 {
				skipCount--
				continue
			}

			for j := range out[finishIdx-len(itemsTypes[i].itemValue) : finishIdx] {
				out[finishIdx-len(itemsTypes[i].itemValue) : finishIdx][j] = itemsTypes[i].itemValue[j]
			}
			finishIdx = finishIdx - len(itemsTypes[i].itemValue)
			out[finishIdx-1] = '/'
			finishIdx--
		}
	}

	if finishIdx == len(out) {
		finishIdx--
		out[finishIdx] = '/'
	} else if out[finishIdx] == 0 {
		out[finishIdx] = '/'
		finishIdx--
	}

	return string(out[finishIdx:])
}
