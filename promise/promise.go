package promise

import (
	"sync"
)

// Promise é a estrutura contendo informações execução futura
type Promise struct {
	cbOk   []func(interface{})
	cbOnce []func(interface{})
	cbErr  []func(string)
	cbDone []func()

	isOk   bool
	isErr  bool
	isDone bool
	isOnce bool

	obj     interface{}
	message string

	mutex *sync.Mutex
}

// New ...
func New() *Promise {
	return &Promise{
		mutex: &sync.Mutex{},
	}
}

func (p *Promise) callOnce(obj interface{}) {
	p.mutex.Lock()
	if p.isOnce {
		p.mutex.Unlock()
		return
	}
	p.isOnce = true
	p.obj = obj
	p.mutex.Unlock()

	for i := 0; i < len(p.cbOnce); i++ {
		p.cbOnce[i](obj)
	}
}

func (p *Promise) callOk(obj interface{}) {
	p.isOk = true
	p.obj = obj

	for i := 0; i < len(p.cbOk); i++ {
		p.cbOk[i](obj)
	}
}

func (p *Promise) callErr(message string) {
	p.isErr = true
	p.message = message

	for i := 0; i < len(p.cbErr); i++ {
		p.cbErr[i](message)
	}
}

func (p *Promise) callDone() {
	p.isDone = true

	for i := 0; i < len(p.cbDone); i++ {
		p.cbDone[i]()
	}
}

// Once ...
func (p *Promise) Once(cb func(interface{})) *Promise {
	p.mutex.Lock()
	p.cbOnce = append(p.cbOnce, cb)
	p.mutex.Unlock()

	return p
}

// Resolve é invokada caso a ação foi executada com sucesso
func (p *Promise) Resolve(obj interface{}) *Promise {
	p.callOnce(obj)
	p.callOk(obj)
	p.callDone()

	return p
}

// Reject é invokada caso a ação foi executada com falha
func (p *Promise) Reject(message string) *Promise {
	p.callErr(message)
	p.callDone()

	return p
}

// Ok recebe callback de espera que será executado caso houve sucesso na ação
func (p *Promise) Ok(cb func(interface{})) *Promise {
	p.mutex.Lock()
	p.cbOk = append(p.cbOk, cb)
	p.mutex.Unlock()

	if p.isOk {
		p.callOk(p.obj)
	}

	return p
}

// Err recebe callback de espera que será executado caso haja falha
func (p *Promise) Err(cb func(string)) *Promise {
	p.mutex.Lock()
	p.cbErr = append(p.cbErr, cb)
	p.mutex.Unlock()

	if p.isErr {
		p.callErr(p.message)
	}

	return p
}

// Done recebe callback que será executado ao final da ação com sucesso ou falha
func (p *Promise) Done(cb func()) *Promise {
	p.mutex.Lock()
	p.cbDone = append(p.cbDone, cb)
	p.mutex.Unlock()

	if p.isDone {
		p.callDone()
	}

	return p
}

// Clear remove todos os callbacks
func (p *Promise) Clear() *Promise {
	p.cbOk = []func(interface{}){}
	p.cbErr = []func(string){}
	p.cbDone = []func(){}

	return p
}
