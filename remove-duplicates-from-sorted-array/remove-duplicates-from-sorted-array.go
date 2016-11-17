package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	ret := []int{}
	for i, v := range nums {
		if i == 0 {
			count++
			ret = append(ret, v)
		} else if nums[i-1] != v {
			count++
			ret = append(ret, v)
		}
	}
	for i, v := range ret {
		nums[i] = v
	}
	return count

}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 4}))
}
