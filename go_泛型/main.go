package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4}
	fmt.Println(reverseInt(in))

	inStr := []string{"1", "a", "c", "d"}
	fmt.Println(reverseStr(inStr))
}

func reverseInt(in []int) []int {
	length := len(in)
	out := make([]int, length)
	fmt.Println(len(out))
	for i := 0; i < length; i++ {
		out[i] = in[length-i-1]
	}
	return out
}

func reverseStr(in []string) []string {
	length := len(in)
	out := make([]string, length)
	fmt.Println(len(out))
	for i := 0; i < length; i++ {
		out[i] = in[length-i-1]
	}
	return out
}

func reverseGenerics[T int|string](in []T) []T {
	length := len(in)
	out := make([]string, length)
	fmt.Println(len(out))
	for i := 0; i < length; i++ {
		out[i] = in[length-i-1]
	}
	return out
}