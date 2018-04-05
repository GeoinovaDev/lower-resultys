package process

import (
	"errors"
	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
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
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return nil, erros.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

func (p *Proc) Close() {
	for {
		err := p.process.Kill()
		if err != nil {
			log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
			continue
		}
		break
	}
}
