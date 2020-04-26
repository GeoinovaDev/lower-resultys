package hook

import "sync"

// Hook ...
type Hook struct {
	mx   *sync.Mutex
	list map[string][]func(...interface{})
}

// New ...
func New() *Hook {
	return &Hook{
		list: make(map[string][]func(...interface{})),
		mx:   &sync.Mutex{},
	}
}

// On ...
func (h *Hook) On(name string, fn func(...interface{})) *Hook {
	h.mx.Lock()
	defer h.mx.Unlock()

	if !h.existName(name) {
		h.list[name] = []func(...interface{}){}
	}

	h.list[name] = append(h.list[name], fn)

	return h
}

// Off ...
func (h *Hook) Off(name string) *Hook {
	h.mx.Lock()
	defer h.mx.Unlock()

	if h.existName(name) {
		h.list[name] = []func(...interface{}){}
	}

	return h
}

// Trigger ...
func (h *Hook) Trigger(name string, params ...interface{}) *Hook {
	h.mx.Lock()
	defer h.mx.Unlock()

	if h.existName(name) {
		for i := 0; i < len(h.list[name]); i++ {
			h.list[name][i](params...)
		}
	}

	return h
}

func (h *Hook) existName(name string) bool {
	if _, ok := h.list[name]; ok {
		return true
	}

	return false
}
