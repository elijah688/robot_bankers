package robot

import (
	"sync"
	"testing"
)

func TestSpendy_Work(t *testing.T) {
	mutex := new(sync.RWMutex)
	cond := sync.NewCond(mutex)
	bankBalance := 100
	spendy := NewSpendy(cond, &bankBalance)
	wg := new(sync.WaitGroup)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			spendy.Work()
		}()

	}
	wg.Wait()

	mutex.RLock()
	if bankBalance != 0 {
		t.Errorf("unexpected bank balance: got %d, want %d", bankBalance, 0)
	}
	mutex.RUnlock()
}
