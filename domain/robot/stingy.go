package robot

import (
	"sync"
)

type Stingy struct {
	cond         *sync.Cond
	bankBallance *int
}

func NewStingy(cond *sync.Cond, bankBallance *int) *Stingy {
	return &Stingy{
		cond,
		bankBallance,
	}
}

func (s *Stingy) Work() {
	s.cond.L.Lock()
	for *s.bankBallance >= 10 {
		s.cond.Wait()
	}
	*s.bankBallance += 10
	s.cond.L.Unlock()
	s.cond.Broadcast()
}
