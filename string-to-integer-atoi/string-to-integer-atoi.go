package main

import (
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	s := ""
	for index, value := range str {
		switch value {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			s += string(value)
			continue
		case '+', '-':
			if index == 0 {
				s += string(value)
				continue
			}
		}
		break
	}
	ret, _ := strconv.Atoi(s)
	if ret >= 2147483648 {
		return 2147483647
	} else if ret < -2147483648 {
		return -2147483648
	}
	return ret
}

func main() {
	fmt.Println(myAtoi("   -5asdf111111a"))
}
