package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup // 等待组
	mutex   sync.Mutex     // 互斥锁
	rwMutex sync.RWMutex   // 读写锁
)

var a map[string]string

// add 对全局变量x执行5000次加1操作ma
func add() {
	mutex.Lock()
	for i := 0; i < 5000; i++ {
		x += 1
	}
	mutex.Unlock()
	wg.Done()
}
func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)

	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000)

	// 使用读写锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000)
	a = make(map[string]string)
	a["hello"] = "world"
}

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10)
	mutex.Unlock()
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

// writeWithRWLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10)
	rwMutex.Unlock()
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()
	time.Sleep(time.Millisecond)
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	// rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x: %v, cost: %vms\n", x, cost.Milliseconds())

}
