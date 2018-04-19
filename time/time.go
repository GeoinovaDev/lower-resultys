package time

import "time"

// Timeout executa o callback depoois de 'second' segundos
func Timeout(second int, callback func()) {
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		callback()
	}()
}
