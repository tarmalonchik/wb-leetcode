package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match(".*aax*v*s*", []byte("abcaaxvs")))
}
