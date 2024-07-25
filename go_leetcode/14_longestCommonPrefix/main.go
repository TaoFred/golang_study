package main

import "fmt"

func main() {
	input := []string{"fegg", "feg", "fedg"}
	output := longtestCommonPrefix(input)
	fmt.Println(output)
}

// 编写一个函数来查找字符串数组中的最长公共前缀
// 如果不存在公共前缀，返回空字符串""
// 	{input: []string{"fegg", "feg", "fedg"}, want: "fe"},
func longtestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for id := 1; id < len(strs); id++ {
		res = lcp(res, strs[id])
		if res == "" {
			return ""
		}
	}
	return res
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	if length == 0 {
		return ""
	}
	id := 0
	for ; id < length; id++ {
		if str1[id] != str2[id] {
			break
		}
	}
	return str1[:id]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
