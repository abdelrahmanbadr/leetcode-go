package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ret := []int{0, 0}
	var n2 []int = nums
	for i1, e1 := range nums {
		for i2, e2 := range n2 {
			if e1+e2 == target {
				if i1 != i2 {
					ret[0] = i1
					ret[1] = i2
					return ret
				}
			}
		}
	}
	return ret
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
