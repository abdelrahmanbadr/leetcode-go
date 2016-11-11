package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return target
	}

	close := 100000000.0
	ret := 0
	for i := 0; len(nums[i:]) > 2; i++ {
		for j := i + 1; len(nums[j:]) > 1; j++ {
			for k := j + 1; len(nums[k:]) > 0; k++ {
				difference := math.Abs(float64(nums[i] + nums[j] + nums[k] - target))
				if difference == 0 {
					return target
				} else if difference < close {
					ret = nums[i] + nums[j] + nums[k]
					close = difference
				}
			}
		}
	}
	return ret

}

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
}
