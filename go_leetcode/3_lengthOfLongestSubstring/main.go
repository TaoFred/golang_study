package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring1("fegjeifejge"))
	fmt.Println(lengthOfLongestSubstring1("abcd"))
	fmt.Println(lengthOfLongestSubstring1("bbbbb"))
	fmt.Println(lengthOfLongestSubstring1(" "))
	fmt.Println(lengthOfLongestSubstring1(""))

	fmt.Println(lengthOfLongestSubstring2("fegjeifejge"))
	fmt.Println(lengthOfLongestSubstring2("abcd"))
	fmt.Println(lengthOfLongestSubstring2("bbbbb"))
	fmt.Println(lengthOfLongestSubstring2(" "))
	fmt.Println(lengthOfLongestSubstring2(""))

}

// 本人错误思路
func lengthOfLongestSubstring(s string) int {
	length := 0
	start := 0
	runeMap := make(map[rune]int, len(s))
	for id, char := range s {
		if valueId, ok := runeMap[char]; ok {
			length = max(length, id-start)
			runeMap[char] = id
			start = valueId + 1
		} else {
			runeMap[char] = id
		}
	}
	length = max(length, len(s)-start) // 很容易缺少这一步，最后需要再比较一次
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口解法
func lengthOfLongestSubstring1(s string) int {
	length := 0
	byteMap := make(map[byte]struct{})

flag:
	for i := 0; i < len(s); i++ {
		for j := i + len(byteMap); j < len(s); j++ {
			if _, ok := byteMap[s[j]]; ok {
				length = max(length, j-i)
				delete(byteMap, s[i])
				continue flag
			} else {
				byteMap[s[j]] = struct{}{}
			}
		}
		length = max(length, len(s)-i)
	}
	return length
}

// 滑动窗口解法2
func lengthOfLongestSubstring2(s string) int {
	length := 0
	rk := 0 // 窗口右边界
	byteMap := make(map[byte]int8)

	for i := 0; i < len(s); i++ {
		for rk < len(s) && byteMap[s[rk]] == 0 {
			// 先赋值再自增，否则索引溢出
			byteMap[s[rk]] = 1
			rk++
		}
		// 此时rk已经是右边界+1了
		length = max(length, rk-i)
		delete(byteMap, s[i])
	}
	return length
}
