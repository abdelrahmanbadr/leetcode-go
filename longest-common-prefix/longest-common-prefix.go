package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ret := ""
	for i, b := range strs[0] {
		for _, s := range strs {
			if len(s) <= i {
				return ret
			}
			if string(s[i]) != string(b) {
				return ret
			}
		}
		ret += string(b)
	}
	return ret

}

func main() {
	strs := []string{"aa", "aa", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
