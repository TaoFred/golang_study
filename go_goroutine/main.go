package main

import "fmt"

func main() {
	test1()
}

func test() {
	fmt.Println("test hello")
}

func test1() {
	fmt.Println("test1 hello")
	go test()
}
