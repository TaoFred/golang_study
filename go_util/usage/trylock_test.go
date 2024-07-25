package usage

import "testing"

func TestTryLock(t *testing.T) {
	for i := 0; i < 10; i++ {
		TryLock()
	}
}
