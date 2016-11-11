package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	if len(nums) <= 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; len(nums[i:]) > 3; i++ {
		for j := i + 1; len(nums[j:]) > 2; j++ {
			for k := j + 1; len(nums[k:]) > 1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] > target {
						break
					}
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						has := false
						for _, r := range ret {
							if r[0] == nums[i] && r[1] == nums[j] && r[2] == nums[k] && r[3] == nums[l] {
								has = true
								break
							}
						}
						if has == false {
							ret = append(ret, []int{nums[i], nums[j], nums[k], nums[l]})
						}
						break
					}
				}

			}
		}
	}
	return ret
}

func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
