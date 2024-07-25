package util_sync

import (
	"sync"
)

// 使用 sync.Once 结构安全地实现go语言单例模式
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

/*
sync.Once结构通过使用原子操作和互斥锁来保证函数只执行一次。具体实现如下：

结构体定义： sync.Once结构体包含一个done字段（类型为uint32）和一个互斥锁m（类型为sync.Mutex）。

Do方法： Do方法首先使用原子操作检查done字段是否已经被设置。如果done字段已经被设置为非零值，表示函数已经执行过，直接返回。

互斥锁： 如果done字段为零，表示函数还没有执行过，Do方法会使用互斥锁来确保只有一个goroutine能够执行函数f。在执行完函数f之后，Do方法会将done字段设置为非零值，表示函数已经执行过。
*/

/*
在并发编程中，slow-path和fast-path是常用的术语，用来描述不同的执行路径：

Fast-path（快速路径）：指的是在大多数情况下执行的代码路径，通常是无锁或低开销的操作。它旨在快速完成任务，以提高性能。

Slow-path（慢速路径）：指的是在某些特殊情况下执行的代码路径，通常涉及锁或其他高开销的操作。它用于处理复杂或罕见的情况，确保正确性。

在sync.Once的实现中：

Fast-path：通过原子操作检查done字段是否已经被设置。如果已经设置，直接返回。这是无锁的快速路径，适用于大多数情况下的检查。

Slow-path：如果done字段为零，表示函数还没有执行过，此时需要进入慢速路径。慢速路径使用互斥锁来确保只有一个goroutine能够执行函数f，并在执行完函数f之后设置done字段。
*/
