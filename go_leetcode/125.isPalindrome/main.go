package main

import (
	"fmt"
	"strings"
)

func main() {
	isPalindrome("abc012")

	res := []int{1}
	fmt.Printf("res: %v\n", res)

	temp := res[1:]
	fmt.Printf("temp: %v\n", temp)

	fmt.Printf("res[1]: %v\n", res[1])
}

func isPalindrome(s string) bool {

	for _, char := range s {
		fmt.Printf("char: %v\n", char)
		fmt.Printf("string(char): %v\n", string(char))
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[i]: %v\n", s[i])
		if s[i] > 'a' {
		}

	}

	s = strings.ToLower(s)

	s = strings.ToLower(s)
	return false
}
