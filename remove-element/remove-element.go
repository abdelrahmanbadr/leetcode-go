package main

import "fmt"

func removeElement(nums []int, val int) int {
	ret := []int{}
	for _, v := range nums {
		if v != val {
			ret = append(ret, v)
		}
	}
	for i, v := range ret {
		nums[i] = v
	}
	return len(ret)

}

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 3, 4}, 3))
}
