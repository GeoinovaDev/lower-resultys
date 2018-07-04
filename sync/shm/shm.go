package shm

import (
	"sync"
)

// Shm struct
type Shm struct {
	wg *sync.WaitGroup
}

// New ...
func New() *Shm {
	return &Shm{
		wg: &sync.WaitGroup{},
	}
}

// Wait ...
func (s *Shm) Wait() {
	s.wg.Add(1)
	s.wg.Wait()
}

// Done ...
func (s *Shm) Done() {
	defer func() {
		recover()
	}()
	s.wg.Done()
}
