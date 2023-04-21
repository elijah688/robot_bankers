package robot

import (
	"sync"
	"testing"
	"time"
)

func TestRobots_Run(t *testing.T) {
	mutex := new(sync.RWMutex)
	bankBalance := 0
	cond := sync.NewCond(mutex)
	robots := NewRobots(cond, &bankBalance)
	changed := false

	// Run the robots
	robots.Run()

	// Check the balance 20 times
	for i := 0; i < 200; i++ {
		// Delay
		time.Sleep(300 * time.Microsecond)

		cond.L.Lock()
		cond.Wait()
		// Log balance change
		if bankBalance > 0 {
			changed = true
		}

		// Check the bank balance oscillate between 0 and 10
		if bankBalance < 0 || bankBalance > 10 {
			t.Fatalf("bank balance should oscillate between 0 and 10, but got %d", bankBalance)
		}
		cond.L.Unlock()

	}
	// Assert that balance changed at least once
	if !changed {
		t.Fatalf("bank balance should have been changed")
	}

}
