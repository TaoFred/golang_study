package main

// func main() {
// 	fmt.Println(pinyin.LazyConvert("中国", &pinyin.Args{}))
// }

// func pinyinConvert(name string) string {
// 	result := pinyin.Pinyin(name, pinyin.Args{Style: pinyin.Normal})
// 	var converted []string
// 	for _, py := range result {
// 		converted = append(converted, py[0])
// 	}
// 	return strings.Join(converted, "")
// }

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	// strSlice := []string{"中", "45", "中", "这", "a", "12"}

	// sort.Slice(strSlice, func(i, j int) bool {
	// 	return less(strSlice[i], strSlice[j])
	// })

	// fmt.Println(strSlice)

	fmt.Println(pinyin.Pinyin("123abd中fefegege国fefeg121", pinyin.NewArgs()))
	printTest()

	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("中"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("美"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("中国大"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("浙  江 大"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("杭 州 333大"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("杭 州大"))
	fmt.Printf("convertPinyin(\"中\"): %v\n", convertPinyin("杭州大"))
}

func getCharType(ch rune) int {
	switch {
	case unicode.IsDigit(ch):
		return 0 // 数字
	case unicode.IsLetter(ch):
		return 1 // 字母
	default:
		return 2 // 中文
	}
}

func less(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		aCharType := getCharType(rune(a[i]))
		bCharType := getCharType(rune(b[i]))

		if aCharType != bCharType {
			return aCharType < bCharType
		} else if aCharType == 2 { // 中文字符按拼音排序
			pinyinA := pinyin.LazyConvert(string(a[i]), nil)
			pinyinB := pinyin.LazyConvert(string(b[i]), nil)
			return strings.Join(pinyinA, "") < strings.Join(pinyinB, "")
		}
	}

	return len(a) < len(b)
}

func printTest() {
	str := "中控supcon1234浙江"
	for _, r := range str {
		fmt.Printf("r: %v\n", r)
	}

	runes := []rune(str)

	for _, r := range runes {
		fmt.Printf("r: %v\n", r)
	}

	for i := 0; i < len(str); i++ {
		r := rune(str[i])
		fmt.Printf("r: %v\n", r)
	}
}

func convertPinyin(hanzi string) string {
	res := pinyin.LazyConvert(hanzi, nil)
	return strings.Join(res, "")
}
