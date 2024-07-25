package usage

import (
	"sync"
	"testing"
)

func TestSingleMachineLock(t *testing.T) {
	for i := 0; i < 10; i++ {
		SingleMachineLock()
	}
}

func TestDistributeLockWithRedis(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DistributeLockWithRedis()
		}()
	}
}
