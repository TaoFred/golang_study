package main

import (
	"fmt"
	"sync"
)

func main() {
	buffer := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		produce(buffer, "producer1")
		wg.Done()
	}()

	for i := 0; i < 2; i++ {
		go func(i int) {
			consume(buffer, fmt.Sprintf("consumer%d", i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func produce(buffer chan<- int, producer string) {
	for i := 1; i < 10; i++ {
		buffer <- i
		fmt.Printf("%s produce task %d\n", producer, i)
	}
	close(buffer)
}

func consume(buffer <-chan int, consumer string) {
	for i := range buffer {
		fmt.Printf("%s consume task %d\n", consumer, i)
	}
}
