package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestSubstringInfo(s string) (int, string) {
	last := make(map[byte]int)
	left := 0
	bestLen := 0
	bestStart := 0

	for right := 0; right < len(s); right++ {
		c := s[right]

		if p, ok := last[c]; ok && p >= left {
			left = p + 1
		}

		last[c] = right

		currentLen := right - left + 1
		if currentLen > bestLen {
			bestLen = currentLen
			bestStart = left
		}
	}

	return bestLen, s[bestStart : bestStart+bestLen]
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

	length, substring := longestSubstringInfo(s)

	fmt.Println("Output:", length)
	fmt.Printf("Explanation: The answer is %q, with the length of %d.\n", substring, length)
}
