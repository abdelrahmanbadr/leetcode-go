package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i := 0
	for j := 0; j < len(needle) && i+j < len(haystack); j++ {
		if haystack[i+j] != needle[j] {
			i++
			j = -1
		} else if j == len(needle)-1 {
			return i
		}

	}
	return -1
}

func main() {
	fmt.Println(strStr("aaaabc", "bc"))
}
