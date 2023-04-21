package robot

import "sync"

type Robot interface {
	Work()
}

type Robots [2]Robot

func NewRobots(condVar *sync.Cond, balance *int) *Robots {
	return &Robots{
		NewSpendy(condVar, balance),
		NewStingy(condVar, balance),
	}
}

func (r *Robots) Run() {
	for _, robot := range r {
		go func(robot Robot) {
			for {
				robot.Work()
			}
		}(robot)
	}
}
