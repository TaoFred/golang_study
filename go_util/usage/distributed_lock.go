package usage

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// 单机锁
/*
定义
单机锁是在单个进程或单个机器上实现的锁机制，用于在同一进程或机器上的多个线程之间进行同步。

实现方式
	互斥锁（Mutex）：使用互斥锁来确保同一时间只有一个线程可以访问共享资源。
	读写锁（RWMutex）：允许多个线程同时读取共享资源，但在写操作时需要独占锁。
适用场景
	适用于单机环境，即所有需要同步的线程都运行在同一台机器上。
	适用于单进程多线程的应用程序。
*/

func SingleMachineLock() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // 加锁，锁的临界区域最小化
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)

}

// 分布式锁
/*
定义
分布式锁是在分布式系统中实现的锁机制，用于在多个进程或多台机器之间进行同步。

实现方式
	基于数据库：使用数据库的行锁或事务机制来实现分布式锁。
	基于缓存（如Redis）：使用Redis的 SETNX 命令来实现分布式锁。
	基于Zookeeper：使用Zookeeper的临时节点和顺序节点来实现分布式锁。
	基于Etcd：使用Etcd的租约机制来实现分布式锁。

适用场景
	适用于分布式环境，即需要在多台机器或多个进程之间进行同步。
	适用于分布式系统中的资源竞争问题。
*/

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:9295",
		Password: "supcon",
	})
)

func init() {
	if rdb.Ping().Err() == nil {
		println("redis client init success")
	} else {
		panic("redis client init failed")
	}
	rdb.Set(counterKey, 0, 0)
}

func acquireLock(key string, value string) bool {
	result, err := rdb.SetNX(key, value, 5*time.Second).Result()
	if err != nil {
		println("setnx failed: %v\n", err)
		return false
	}
	println("setnx success")
	return result
}

func releaseLock(key string) bool {
	delResp := rdb.Del(key)
	del, err := delResp.Result()
	if err == nil && del > 0 {
		println("unlock success")
		return true
	}
	println("unlock failed")
	return false
}

var lockKey = "counter_lock"
var counterKey = "counter"

func CounterIncr() {
	getResp := rdb.Get(counterKey)
	counterVal, err := getResp.Int()
	if err != nil {
		println("get counter value err: ", err)
		return
	}
	counterVal++
	setResp := rdb.Set(counterKey, counterVal, 0)
	_, err = setResp.Result()
	if err != nil {
		println("set counter value err: ", err)
	}
	println("current counter is", counterVal)

}

func DistributeLockWithRedis() {
	if !acquireLock(lockKey, "counter_lock_val") {
		return
	}
	CounterIncr()
	releaseLock(lockKey)
}
