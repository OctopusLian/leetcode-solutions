package main

import (
	"fmt"
	"strings"
)

func main() {
	if isValid("()") {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	for len(s) != 0 {
		temp := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)

		if s == temp {
			return false
		}
	}

	return true
}
