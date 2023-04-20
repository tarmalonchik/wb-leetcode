package main

import (
	"fmt"
)

const (
	symTrue  = "("
	symFalse = ")"
)

func main() {
	num := 3
	fmt.Println(generateParenthesis(num))
}

func generateParenthesis(n int) []string {
	rootNode := nodeType{
		depth:        n * 2,
		possibleZero: 1,
		possibleOne:  n - 1,
		symbolString: symTrue,
	}
	var resp totalStrings
	resp.fillNode(&rootNode)
	return resp
}

type totalStrings []string

type nodeType struct {
	depth        int
	possibleZero int
	possibleOne  int
	rNode        *nodeType
	lNode        *nodeType
	symbolString string
}

func (t *totalStrings) fillNode(node *nodeType) {
	if node == nil {
		return
	}
	if node.depth == 1 {
		*t = append(*t, node.symbolString)
		return
	}
	if node.possibleZero > 0 {
		node.rNode = &nodeType{
			depth:        node.depth - 1,
			possibleZero: node.possibleZero - 1,
			possibleOne:  node.possibleOne,
			symbolString: node.symbolString + symFalse,
		}
		t.fillNode(node.rNode)
	}
	if node.possibleOne > 0 {
		node.lNode = &nodeType{
			depth:        node.depth - 1,
			possibleZero: node.possibleZero + 1,
			possibleOne:  node.possibleOne - 1,
			symbolString: node.symbolString + symTrue,
		}
		t.fillNode(node.lNode)
	}
}
