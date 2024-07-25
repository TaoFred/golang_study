package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	exist := make(map[[3]int]bool)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				vi := nums[i]
				vj := nums[j]
				vk := nums[k]
				if vi+vj+vk == 0 && !exist[[3]int{vi, vj, vk}] {
					res = append(res, []int{vi, vj, vk})
					exist[[3]int{vi, vj, vk}] = true
				}
			}
		}

	}

	return res
}
