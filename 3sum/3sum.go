package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; len(nums[i:]) > 2; i++ {
		j := i + 1
		for len(nums[j:]) > 1 {
			k := len(nums) - 1
			for nums[i]+nums[j]+nums[k] > 0 && k > j {
				k--
			}
			if k > j {
				if nums[i]+nums[j]+nums[k] == 0 {
					has := false
					for _, r := range ret {
						if r[0] == nums[i] && r[1] == nums[j] && r[2] == nums[k] {
							has = true
							break
						}
					}
					if has == false {
						ret = append(ret, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
			j++
		}
	}
	return ret

}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
