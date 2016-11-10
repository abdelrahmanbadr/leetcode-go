package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		now := i
		interval := 2 * i
		for now < len(s) {
			ret += string(s[now])
			if i == 0 || i == numRows-1 {
				interval = 2*numRows - 2
			} else {
				if interval == 2*numRows-2*i-2 {
					interval = 2 * i
				} else {
					interval = 2*numRows - 2*i - 2
				}
			}

			now += interval
		}
	}
	return ret
}

func main() {
	s := "PAYPALISHIRING"
	r := 3
	fmt.Println(convert(s, r))
}
