package process

import (
	"git.resultys.com.br/framework/lower/exception"
	"os"
	"os/exec"
)

type Proc struct {
	process *os.Process
	state   *os.ProcessState
}

func Start(program string) (proc *Proc, err *exception.Error) {
	cmd := exec.Command(program)
	err1 := cmd.Start()
	if err1 != nil {
		return nil, &exception.Error{What: "processo nao encontrado"}
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

func (p *Proc) Close() {
	for {
		err := p.process.Kill()
		if err != nil {
			continue
		}
		break
	}
}
