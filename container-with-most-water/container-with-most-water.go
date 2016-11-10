package main

import "fmt"

func maxArea(height []int) int {
	ret := 0
	for i, a := range height {
		for j := i + 1; j < len(height); j++ {
			h := height[j]
			if a < h {
				h = a
			}
			v := (j - i) * h
			if v > ret {
				ret = v
			}
		}
	}
	return ret
}

func main() {
	height := []int{1, 1, 3, 5, 6, 7, 9, 2, 3, 5}
	fmt.Println(maxArea(height))
}
