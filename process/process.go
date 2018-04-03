package process

import (
	"errors"
	"os"
	"os/exec"
)

type Proc struct {
	process *os.Process
	state   *os.ProcessState
}

func Start(program string) (proc *Proc, err error) {
	cmd := exec.Command(program)
	err1 := cmd.Start()
	if err1 != nil {
		return nil, erros.New("processo nao encontrado")
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
