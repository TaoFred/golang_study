package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(nums))
}

// 给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
func findShortestSubArray(nums []int) int {
	type item struct {
		count   int
		fisrtId int
		lastId  int
	}

	numCountMap := make(map[int]*item)
	for id, num := range nums {
		if _, ok := numCountMap[num]; ok {
			numCountMap[num].count++
			numCountMap[num].lastId = id
		} else {
			numCountMap[num] = &item{
				count:   1,
				fisrtId: id,
				lastId:  id,
			}
		}
	}

	minInterval := len(nums)
	maxCount := 0
	for _, item := range numCountMap {
		if item.count > maxCount {
			maxCount = item.count
			minInterval = (item.lastId - item.fisrtId + 1)
		} else if item.count == maxCount {
			interval := (item.lastId - item.fisrtId + 1)
			if minInterval > interval {
				minInterval = interval
			}
		}
	}
	return minInterval
}
