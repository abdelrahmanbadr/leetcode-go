package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0.0
	}
	nums := make([]int, m+n, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			k++
			i++
		} else {
			nums[k] = nums2[j]
			k++
			j++
		}
	}
	if i < m {
		nums = append(nums[:k], nums1[i:]...)
	} else {
		nums = append(nums[:k], nums2[j:]...)
	}

	for _, e := range nums {
		fmt.Println(e)
	}

	ret := (m + n) / 2
	if (m+n)%2 == 1 {
		return float64(nums[ret])
	} else {
		return float64(float64((nums[ret] + nums[ret-1])) / 2.0)
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ret := 0.0
	ret = findMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)
}
