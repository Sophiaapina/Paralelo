package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(s string) bool {

	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {

		// si es apertura → push
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		// si es cierre
		if open, ok := pairs[c]; ok {

			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != open {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input: s = ")
	s, _ := reader.ReadString('\n')

	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '\r' {
		s = s[:len(s)-1]
	}

	result := isValid(s)

	fmt.Println("Output:", result)
}
