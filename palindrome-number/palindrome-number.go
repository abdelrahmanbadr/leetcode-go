package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for i, v := range s {
		if i > l/2 {
			break
		}
		if s[l-i-1] != byte(v) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(312))
}
