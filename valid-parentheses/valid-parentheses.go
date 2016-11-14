package main

import "fmt"

func isValid(s string) bool {
	stack := []string{}
	for _, c := range s {
		l := len(stack)
		switch string(c) {
		case "(", "[", "{":
			stack = append(stack, string(c))
		case ")":
			if l > 0 && string(stack[l-1]) == "(" {
				stack = stack[:l-1]
			} else {
				return false
			}
		case "]":
			if l > 0 && string(stack[l-1]) == "[" {
				stack = stack[:l-1]
			} else {
				return false
			}
		case "}":
			if l > 0 && string(stack[l-1]) == "{" {
				stack = stack[:l-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("]"))
}
