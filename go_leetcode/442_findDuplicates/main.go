package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates1(nums))
}

func findDuplicates(nums []int) (ans []int) {
	tempMap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := tempMap[num]; ok {
			ans = append(ans, num)
		} else {
			tempMap[num] = struct{}{}
		}
	}
	return
}

func findDuplicates1(nums []int) (ans []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return
}
