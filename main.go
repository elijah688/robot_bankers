package main

import (
	"fmt"
    "github.com/elijah688/robot_bankers/domain/robot"
	"sync"
	"time"
)

var (
	mutex        = new(sync.RWMutex)
	bankBallance = new(int)
	cond         = sync.NewCond(mutex)
)

func main() {
	robot.NewRobots(cond, bankBallance).Run()
	for {
		time.Sleep(200 * time.Millisecond)
		mutex.RLock()
		fmt.Println(*bankBallance)
		mutex.RUnlock()
	}

}
