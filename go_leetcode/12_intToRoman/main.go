package main

import "fmt"

func main() {

	// inputInt := 1993
	// roman := intToRoman(inputInt)
	// fmt.Println(roman)
	// fmt.Println(romanToInt(roman))
	fmt.Println(romanToInt("III"))
}

func test(x, y int) (int, int) {
	shang := x / y
	yu := x % y

	return shang, yu
}

var numLabel = []struct {
	num   int
	lable string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, nl := range numLabel {
		for num >= nl.num {
			roman = append(roman, nl.lable...)
			num -= nl.num
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

var lableNum = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	for i := range s {
		value := lableNum[s[i]]
		// 在对索引操作时，要特别注意越界问题
		if i+1 < len(s) && lableNum[s[i+1]] > value {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}
