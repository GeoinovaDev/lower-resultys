package time

import (
	"time"
)

type timeout struct {
	isTrigger bool
}

func (t *timeout) Clear() {
	t.isTrigger = false
}

// Timeout executa o callback depoois de 'second' segundos
func Timeout(second int, callback func()) *timeout {
	t := &timeout{isTrigger: true}
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		if t.isTrigger {
			callback()
		}
	}()

	return t
}
