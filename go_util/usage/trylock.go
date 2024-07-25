package usage

import (
	"fmt"
	"sync"
)

// trylock
// TryLock 是一种非阻塞锁机制，它尝试获取锁，
// 如果锁已经被其他线程或进程持有，则立即返回失败，而不是等待锁释放。

// 使用大小为1的channel实现trylock
type tryLock struct {
	c chan struct{}
}

func NewTryLock() tryLock {
	var l tryLock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l tryLock) Lock() bool {
	select {
	case <-l.c:
		return true
	default:
		return false
	}
}

func (l tryLock) Unlock() {
	l.c <- struct{}{}
}

func TryLock() {
	l := NewTryLock()
	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if l.Lock() {
				counter++
				// fmt.Println("current counter", counter)
				l.Unlock()
				return
			}
			// fmt.Println("failed to get lock")
		}()
	}
	wg.Wait()
	fmt.Println("final counter", counter)
}
