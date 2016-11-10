package main

import "fmt"

func charInString(a string, s string) (bool, int){
	for i, len := 0, len(s); i < len; i++{
		if a == string(s[i]){
			return true, i
		}
	}
	return false, len(s) 
}

func lengthOfLongestSubstring(s string) int {
	l := ""
	ret := 0
	for i, length := 0, len(s); i < length; i++{
		su, index := charInString(string(s[i]), l)
		if su{
			if ret < len(l){
				ret = len(l)
			}
			l = string(l[index+1:])
			l += string(s[i])
		}else{
			l += string(s[i])
		}
	} 
	if ret < len(l){
		ret = len(l)
	}
	return ret
}

func main() {
	ret := 0
	ret = lengthOfLongestSubstring("aaasdfas212syb")
	fmt.Println(ret)
}
