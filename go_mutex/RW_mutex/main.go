package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 安管平台事件类型
var SAFETY_MANAGE_EVENT_TYPE = []SafetyManageEventType{
	{
		Key:   201,
		Value: "暴力破解",
	},
	{
		Key:   3000,
		Value: "异常关机",
	},
	{
		Key:   3001,
		Value: "程序异常拦截",
	},
	{
		Key:   3002,
		Value: "异常访问",
	},
	{
		Key:   3003,
		Value: "程序高危拦截",
	},
	{
		Key:   3004,
		Value: "主机资源异常",
	},
	{
		Key:   3005,
		Value: "主机安全卫士防护停止",
	},
	{
		Key:   3006,
		Value: "高危网络拦截",
	},
	{
		Key:   3007,
		Value: "配置修改",
	},
	{
		Key:   3008,
		Value: "异常网络拦截",
	},
	{
		Key:   3009,
		Value: "发现病毒",
	},
	{
		Key:   9000,
		Value: "其他",
	},
}

// 事件类型的读写锁
var EVENT_TYPE_MUTEX sync.RWMutex

// 事件类型
type SafetyManageEventType struct {
	Key   int    `json:"name"`  // 整型枚举
	Value string `json:"value"` // 事件类型名称
}

// 写入数据
func ReadSafetyManageEventType() {
	// EVENT_TYPE_MUTEX.RLock()
	// defer EVENT_TYPE_MUTEX.RUnlock()
	for _, item := range SAFETY_MANAGE_EVENT_TYPE {
		fmt.Println(item.Key)
		fmt.Println(item.Value)
	}
}

// 更新数据
func UpdateSafetyManageEventType(key1, key2 int) {
	// EVENT_TYPE_MUTEX.Lock()
	// defer EVENT_TYPE_MUTEX.Unlock()
	SAFETY_MANAGE_EVENT_TYPE = []SafetyManageEventType{
		{
			Key:   key1,
			Value: "自定义" + strconv.Itoa(key1),
		},
		{
			Key:   key2,
			Value: "自定义" + strconv.Itoa(key2),
		},
	}
}

// 定期更新数据
func periodUpdateData() {
	periodTime := time.NewTicker(10 * time.Second)
	defer periodTime.Stop()
	i := 0
	for range periodTime.C {
		i++
		UpdateSafetyManageEventType(i, i+1)
	}
}

// 定期更新数据
func periodReadData() {
	periodTime := time.NewTicker(5 * time.Second)
	defer periodTime.Stop()
	for range periodTime.C {
		ReadSafetyManageEventType()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		periodReadData()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		periodUpdateData()
		wg.Done()
	}()

	wg.Wait()
}
