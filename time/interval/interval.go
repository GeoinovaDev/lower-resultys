package interval

import (
	"time"
)

// Interval ...
type Interval struct {
	cycle    bool
	overflow bool
}

// New ...
func New() *Interval {
	return &Interval{
		cycle:    true,
		overflow: true,
	}
}

// Clear ...
func (t *Interval) Clear() {
	t.cycle = false
	t.overflow = false
}

// Repeat ...
func (t *Interval) Repeat(second int, callback func()) *Interval {
	go func() {
		for {
			time.Sleep(time.Duration(second) * time.Second)
			if !t.cycle {
				break
			}

			callback()
		}
	}()

	return t
}

// Timeout executa o callback depoois de 'second' segundos
func (t *Interval) Timeout(second int, callback func()) *Interval {
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		if t.overflow {
			callback()
		}
	}()

	return t
}
