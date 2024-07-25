package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers1(nums))
}

func findDisappearedNumbers(nums []int) (ans []int) {
	for i := 1; i < len(nums)+1; i++ {
		if !isExist(nums, i) {
			ans = append(ans, i)
		}
	}
	return
}

func isExist(nums []int, i int) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}

func findDisappearedNumbers1(nums []int) (ans []int) {
	tempMap := make(map[int]struct{})
	for _, num := range nums {
		tempMap[num] = struct{}{}
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := tempMap[i]; !ok {
			ans = append(ans, i)
		}
	}
	return
}
