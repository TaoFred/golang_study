package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

type testStruct struct {
	id   int
	name string // 对此字段进行排序
}

var tests = []testStruct{
	{1, "response"},
	{2, "文档管理系统调研"},
	{3, "概要设计说明书-文档仓库"},
	{4, "文档仓库关联工控资产"},
	{5, "概要设计说明书-全局搜索 - 副本"},
	{6, "概要设计说明书-全局搜索"},
	{7, "图片2"},
	{8, "图片1"},
	{9, "minio"},
	{10, "演示文稿1"},
	{11, "组件结构图"},
	{12, "安管平台"},
	{13, "文档仓库结构"},
	{14, "filebrowser"},
	{15, "minio-server后台服务"},
	{16, "ECS700"},
	{17, "mino-server"},
	{18, "minio控制台"},
	{19, "端口分配_2023-07-27"},
}

func main() {
	sort.Slice(tests, func(i, j int) bool {
		return CompareObjectByStr(tests[i].name, tests[j].name)
	})

	fmt.Println(tests)
}

// 对包含字符串字段的切片进行排序
// 排序规则:
// 特殊字符（非数字，字母，汉字）> 数字 > 字母 > 中文(中文按拼音排序)
// 结合Sort.Slice()使用
func CompareObjectByStr(nameA, nameB string) bool {
	// 字母排序不区分大小写
	runeSliceA := []rune(strings.ToLower(nameA))
	runeSliceB := []rune(strings.ToLower(nameB))
	mixLen := len(runeSliceA)
	if mixLen > len(runeSliceB) {
		mixLen = len(runeSliceB)
	}

	for i := 0; i < mixLen; i++ {
		runeA := runeSliceA[i]
		runeB := runeSliceB[i]

		switch {
		case isSpecialChar(runeA) && !isSpecialChar(runeB):
			return true
		case !isSpecialChar(runeA) && isSpecialChar(runeB):
			return false
		case isDigit(runeA) && !isDigit(runeB):
			return true
		case !isDigit(runeA) && isDigit(runeB):
			return false
		case isLetter(runeA) && !isLetter(runeB):
			return true
		case !isLetter(runeA) && isLetter(runeB):
			return false
		case isChinese(runeA) && !isChinese(runeB):
			return false
		case !isChinese(runeA) && isChinese(runeB):
			return true
		case isChinese(runeA) && isChinese(runeB):
			pinyinA := convertPinyin(string(runeA))
			pinyinB := convertPinyin(string(runeB))
			if pinyinA != pinyinB {
				return pinyinA < pinyinB
			}
		case runeA < runeB:
			return true
		case runeA > runeB:
			return false
		}
	}
	return len(nameA) < len(nameB)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func isChinese(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

func isSpecialChar(r rune) bool {
	return !isDigit(r) && !isLetter(r) && !isChinese(r)
}

func convertPinyin(hanzi string) string {
	res := pinyin.LazyConvert(hanzi, nil)
	return strings.Join(res, "")
}
