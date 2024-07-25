package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	PATH      = "./conf/"
	FILENAME  = "自定义典型回路导入初步模板.xlsx"
	SHEETNAME = "形态识别初步模板"
)

func main() {
	filename := PATH + FILENAME
	xlsxFile, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	rows, err := xlsxFile.GetRows(SHEETNAME)
	if err != nil {
		panic(err)
	}

	cols, err := xlsxFile.GetCols(SHEETNAME)
	if err != nil {
		panic(err)
	}

	fmt.Println("行数据")
	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println("列数据")
	splitId := make([]int, 0)
	for id, col := range cols[6][2:] {
		if col == "" {
			splitId = append(splitId, id)
		}
	}
	nodeInfoTypeList := SplitStringSliceByIndex(cols[2][2:], splitId)
	nodeInfoNameList := SplitStringSliceByIndex(cols[3][2:], splitId)
	beforeNodeNameList := SplitStringSliceByIndex(cols[4][2:], splitId)
	beforeNodePinList := SplitStringSliceByIndex(cols[5][2:], splitId)
	pinRelationList := SplitStringSliceByIndex(cols[6][2:], splitId)
	afterNodeNameList := SplitStringSliceByIndex(cols[7][2:], splitId)
	afterNodePinList := SplitStringSliceByIndex(cols[8][2:], splitId)

	fmt.Println(nodeInfoTypeList, nodeInfoNameList, beforeNodeNameList, beforeNodePinList, pinRelationList, afterNodeNameList, afterNodePinList)
	// isDuplicate := hasDuplicate(cols[3][2:])
	// fmt.Println(isDuplicate)
	// test()
}

// 判断字符串切片是否存在重复元素
func hasDuplicate(s []string) bool {
	m := make(map[string]bool) // 创建一个空的map
	for _, v := range s {      // 遍历切片
		if m[v] { // 如果map中已经存在该元素，说明有重复
			return true
		}
		m[v] = true // 否则，将该元素作为键存入map中
	}
	return false // 遍历完毕，没有发现重复，返回false
}

// 定义一个正则表达式，匹配ipv4地址的格式
var ipv4Regexp = regexp.MustCompile(`^(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)$`)

// 定义一个函数，验证ipv4地址是否合法，并输出其4部分的信息
func checkIPv4(ip string) {
	if ipv4Regexp.MatchString(ip) { // 如果匹配成功，说明是合法的ipv4地址
		fmt.Println("这是一个合法的ipv4地址")
		parts := strings.Split(ip, ".") // 使用"."作为分隔符，将地址分割成四部分
		for i, part := range parts {
			fmt.Printf("第%d部分是：%s\n", i+1, part) // 输出每一部分的信息
		}
	} else { // 如果匹配失败，说明不是合法的ipv4地址
		fmt.Println("这不是一个合法的ipv4地址")
	}
}

func test() {
	checkIPv4("192.168.1.1")   // 测试一个合法的ipv4地址
	checkIPv4("256.100.50.10") // 测试一个不合法的ipv4地址
}

func SplitStringSlice(s []string) [][]string {
	retList := make([][]string, 0)
	ret := make([]string, 0)
	for _, v := range s {
		if v != "" {
			ret = append(ret, v)
			continue
		}
		if len(ret) != 0 {
			retList = append(retList, ret)
		}
		ret = []string{}
	}
	if len(ret) != 0 {
		retList = append(retList, ret)
	}
	return retList
}

func SplitStringSliceByIndex(s []string, splitIds []int) [][]string {
	retList := make([][]string, 0)
	if len(splitIds) == 0 {
		retList = append(retList, s)
		return retList
	}
	retList = append(retList, s[:splitIds[0]])
	for id := 1; id < len(splitIds); id++ {
		ret := s[splitIds[id-1]+1 : splitIds[id]]
		retList = append(retList, ret)
	}
	retList = append(retList, s[splitIds[len(splitIds)-1]+1:])
	return retList
}
