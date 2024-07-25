package main

import (
	"fmt"
	"sync"
)

func main() {
	buffer1 := make(chan int, 10)
	buffer2 := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(buffer1, "producer1")
		wg.Done()
	}()

	go func() {
		produce(buffer2, "producer2")
		wg.Done()
	}()

	select {

	case x, ok := <-buffer1:
		if !ok {
			return
		}
		fmt.Printf("consume producer1 task: %d\n", x)

	case y, ok := <-buffer2:
		if !ok {
			return
		}
		fmt.Printf("consume producer2 task: %d\n", y)
	}

	wg.Wait()
}

func produce(buffer chan<- int, producer string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s produce task: %d\n", producer, i)
		buffer <- i
	}
	close(buffer)
}

func consume(buffer <-chan int) {

	/*
		多返回值判断通道是否关闭
	*/
	// for {
	// 	v, ok := <-buffer
	// 	if !ok {
	// 		fmt.Println("chan is close")
	// 		break
	// 	}
	// 	fmt.Printf("consume task: %d\n", v)
	// }

	/*
		for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环
	*/
	for i := range buffer {
		fmt.Printf("consume task: %d\n", i)
	}
	fmt.Println("chan is close")
}
