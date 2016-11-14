package main

import "fmt"

func findString(ret *[]string, p string, deep, n, left, right int) {
	if deep == 2*n {
		*ret = append(*ret, p)
		return
	}
	if left < n {
		p += "("
		findString(ret, p, deep+1, n, left+1, right)
		p = string(p[:len(p)-1])
	}
	if right < left {
		p += ")"
		findString(ret, p, deep+1, n, left, right+1)
		p = string(p[:len(p)-1])
	}
	return
}

func generateParenthesis(n int) []string {
	ret := []string{}
	if n != 0 {
		findString(&ret, "", 0, n, 0, 0)
	}
	return ret
}

func main() {
	fmt.Println(generateParenthesis(3))
}
