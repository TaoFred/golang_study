package usage

import (
	"sync"
	"testing"
	"time"
)

func TestSingleflight(t *testing.T) {
	var wg sync.WaitGroup
	stratTime := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			getData("key")
		}(i)
	}
	wg.Wait()
	t.Logf("time cost: %v", time.Since(stratTime))
}
