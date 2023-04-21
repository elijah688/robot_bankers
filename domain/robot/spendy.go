package robot

import (
	"sync"
)

type Spendy struct {
	cond         *sync.Cond
	bankBallance *int
}

func NewSpendy(cond *sync.Cond, bankBallance *int) *Spendy {
	return &Spendy{
		cond,
		bankBallance,
	}
}

func (s *Spendy) Work() {
	s.cond.L.Lock()
	for *s.bankBallance < 10 {
		s.cond.Wait()
	}
	*s.bankBallance -= 10
	s.cond.L.Unlock()
	s.cond.Broadcast()
}
