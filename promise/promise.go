package promise

// Promise é a estrutura contendo informações execução futura
type Promise struct {
	cbOk   []func(interface{})
	cbErr  []func(string)
	cbDone []func()

	isOk   bool
	isErr  bool
	isDone bool

	obj     interface{}
	message string
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

// Resolve é invokada caso a ação foi executada com sucesso
func (p *Promise) Resolve(obj interface{}) {
	p.callOk(obj)
	p.callDone()
}

// Reject é invokada caso a ação foi executada com falha
func (p *Promise) Reject(message string) {
	p.callErr(message)
	p.callDone()
}

// Ok recebe callback de espera que será executado caso houve sucesso na ação
func (p *Promise) Ok(cb func(interface{})) *Promise {
	p.cbOk = append(p.cbOk, cb)

	if p.isOk {
		p.callOk(p.obj)
	}

	return p
}

// Err recebe callback de espera que será executado caso haja falha
func (p *Promise) Err(cb func(string)) *Promise {
	p.cbErr = append(p.cbErr, cb)

	if p.isErr {
		p.callErr(p.message)
	}

	return p
}

// Done recebe callback que será executado ao final da ação com sucesso ou falha
func (p *Promise) Done(cb func()) *Promise {
	p.cbDone = append(p.cbDone, cb)

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
