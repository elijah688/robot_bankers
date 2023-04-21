package robot

import (
	"sync"
	"testing"
)

func Test_StingyWork(t *testing.T) {
	mutex := new(sync.RWMutex)
	cond := sync.NewCond(mutex)
	bankBalance := 0
	stingy := NewStingy(cond, &bankBalance)
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			stingy.Work()
		}()
	}
	wg.Wait()

	mutex.RLock()
	if bankBalance != 10 {
		t.Errorf("unexpected bank balance: got %d, want %d", bankBalance, 10)
	}
	mutex.RUnlock()
}
