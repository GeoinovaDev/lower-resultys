package exception

import (
	"time"
)

type Error struct {
	What string
	When time.Time
}

func (e *Error) Error() string {
	return e.What
}
