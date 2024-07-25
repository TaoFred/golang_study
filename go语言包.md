## `sync/atomic`原子操作

原子操作是一种在多线程或多协程环境中安全执行的操作，它保证了在执行过程中不会被其他线程或协程中断。在Go语言中，`sync/atomic`包提供了一系列用于原子操作的函数，这些操作主要用于管理共享状态的同步访问，避免在并发访问时出现竞态条件。

### 使用场景

1. **共享计数器**：在多个goroutine中共享的计数器，比如在线用户数、完成任务的数量等。
2. **状态标志**：表示某个状态是否被设置，如示例中的`isAuth`变量用于标识授权状态。
3. **配置更新**：在不同goroutine中共享的配置信息，需要确保读写操作的原子性。

### 方法

- **读取操作**：`atomic.LoadInt32(&value)`安全地读取`int32`类型的变量`value`。
- **写入操作**：`atomic.StoreInt32(&value, newValue)`安全地将`newValue`写入`int32`类型的变量`value`。
- **增减操作**：`atomic.AddInt32(&value, delta)`将`delta`加到`value`上，并返回新值。`delta`可以是负数，实现减法操作。
- **比较并交换**：`atomic.CompareAndSwapInt32(&value, old, new)`如果当前值等于`old`值，则将`value`更新为`new`值，并返回`true`；否则，不做任何操作并返回`false`。

```go
var isAuth int32 = 0

// 确认是否授权
func (h *BizHandler) IsShellAuth() bool {
	return atomic.LoadInt32(&isAuth) == 1
}

// 周期性检查授权
func (h *BizHandler) PeriodCheckAuth() {
	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()
	for range ticker.C {
		authed := h.HasNeuroShellForEAMAuth()
		if authed && atomic.CompareAndSwapInt32(&isAuth, 0, 1) {
			log.ZapLog.Info("NeuroShellForEAM has been authorized")
		} else if !authed && atomic.CompareAndSwapInt32(&isAuth, 1, 0) {
			log.ZapLog.Info("NeuroShellForEAM auth failed")
		}
	}
}
```



### 示例解析

在提供的`auth.go`文件片段中，使用了原子操作来安全地更新授权状态：

- 使用`atomic.LoadInt32(&isAuth)`安全地读取`isAuth`的值，以判断当前是否已授权。
- 使用`atomic.CompareAndSwapInt32(&isAuth, 0, 1)`和`atomic.CompareAndSwapInt32(&isAuth, 1, 0)`在状态变化时安全地更新`isAuth`的值。这两个操作确保了在多个goroutine尝试修改`isAuth`值时的线程安全性。

通过这种方式，即使`PeriodCheckAuth`方法在多个goroutine中并发执行，`isAuth`的读取和更新也是安全的，避免了竞态条件的发生。这是原子操作在并发编程中非常典型的应用场景。



## 条件变量sync.Cond

> 条件变量是与互斥锁结合使用的，用于在一些条件成立之前阻塞一组goroutine，并在条件成立时唤醒它们。条件变量属于`sync`包，通过`sync.Cond`类型提供



### 使用场景

`sync.Cond`（条件变量）在Go语言中主要用于协调需要等待特定条件满足的goroutine之间的同步。条件变量应该与互斥锁（`sync.Mutex`）一起使用，以避免竞态条件。使用场景包括，但不限于：

1. **等待资源可用**：当goroutine需要等待某种资源（例如，缓冲区中的空间）变得可用时。
2. **生产者-消费者问题**：在生产者和消费者之间同步，其中消费者等待生产者生产新的数据。
3. **等待任务完成**：等待一个或多个并发操作完成，例如，当一个goroutine需要等待另一个goroutine完成某个任务后才能继续执行。

### 使用方法

使用`sync.Cond`的基本步骤如下：

1. **创建条件变量**：使用`sync.NewCond`并传入一个`sync.Locker`（通常是`sync.Mutex`或`sync.RWMutex`）来创建一个新的条件变量。
2. **等待条件满足**：在goroutine中，首先锁定互斥锁，然后调用`Cond.Wait()`方法等待条件满足。`Wait`方法会自动释放锁并阻塞调用的goroutine，直到其他goroutine在相同的条件变量上调用`Signal`或`Broadcast`。
3. **更改条件并通知等待的goroutine**：修改使条件满足的数据后，可以调用`Cond.Signal()`唤醒一个等待的goroutine，或者调用`Cond.Broadcast()`唤醒所有等待的goroutine。
4. **重新检查条件**：被`Signal`或`Broadcast`唤醒的goroutine会重新获取锁，并应该重新检查条件是否真的满足，因为`Signal`和`Broadcast`可能会误唤醒goroutine。

注意，`Wait`方法在阻塞goroutine之前会自动释放锁，并在重新唤醒goroutine时再次获取锁。这是为了确保在等待期间其他goroutine可以获取锁并在条件变量上调用`Signal`或`Broadcast`方法。