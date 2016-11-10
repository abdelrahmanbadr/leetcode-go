package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	sub_str := ""
	if len(s) < 3 {
		return s
	}
	sub_str = string(s[0])
	for i := 0; i < len(s); i++ {
		l := len(sub_str)
		for j := i + l; j < len(s); j++ {
			if s[i] == s[j] {
				if isPalindrome(string(s[i : j+1])) {
					if len(string(s[i:j+1])) >= len(sub_str) {
						sub_str = string(s[i : j+1])
					}
				}
			}
		}
	}
	return sub_str
}

func main() {
	s := "abb"
	fmt.Println(longestPalindrome(s))
}
