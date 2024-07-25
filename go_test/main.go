package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Printf("slice: %v\n", slice)
	appendSlice(slice)
	fmt.Printf("slice: %v\n", slice)
}

func appendSlice(slice []string) {
	slice = slice[:0]
}
