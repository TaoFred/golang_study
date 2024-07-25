package usage

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
singleflight 是 Go 语言标准库 golang.org/x/sync/singleflight 中的一个包，用于解决重复请求的问题。
它的主要功能是确保在高并发环境下，对于相同的请求，只会执行一次操作，其他请求会等待这次操作的结果，从而避免重复计算或重复请求。

使用场景
缓存穿透：在缓存系统中，如果某个缓存失效，多个请求同时到达，会导致多次对数据库或其他后端服务的请求。使用 singleflight 可以确保只有一个请求会去查询数据库，其他请求等待结果返回。
防止重复计算：在高并发环境下，某些计算可能非常耗时且结果相同。使用 singleflight 可以确保同一时间只有一个计算在进行，其他请求复用计算结果。
限流：在某些情况下，限制对某些资源的访问频率，确保不会因为高并发导致资源过载。
*/

var (
	cache             = make(map[string]string)
	mu                sync.Mutex
	singleflightGroup singleflight.Group
)

func fetchFromDB(key string) (interface{}, error) {
	fmt.Printf("fetch %s from db...\n", key)
	time.Sleep(1 * time.Second) // 模拟查询数据库
	return key + "bar", nil
}

func getData(key string) (string, error) {
	if value, ok := cache[key]; ok {
		return value, nil
	}
	v, err, _ := singleflightGroup.Do(key, func() (interface{}, error) {
		// 再次查询缓存，避免重复查询数据库
		if value, ok := cache[key]; ok {
			return value, nil
		}
		value, err := fetchFromDB(key)
		if err != nil {
			return nil, err
		}
		// 更新缓存
		mu.Lock()
		cache[key] = value.(string)
		mu.Unlock()
		return value, nil
	})
	if err != nil {
		return "", err
	}
	return v.(string), nil
}
