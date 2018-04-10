package process

import (
	"errors"
	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Proc struct {
	process *os.Process
	state   *os.ProcessState
}

func Command(command string) (*Proc, error) {
	if runtime.GOOS == "windows" {
		return Cmd(command)
	} else {
		return Bash(command)
	}
}

func Bash(command string) (*Proc, error) {
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

func Cmd(command string) (*Proc, error) {
	cmd := exec.Command("cmd", "/c", command)
	err := cmd.Run()
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

func Start(program string) (proc *Proc, err error) {
	cmd := exec.Command(program)
	err1 := cmd.Start()
	if err1 != nil {
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

func (p *Proc) GetId() int {
	return p.process.Pid
}

func (p *Proc) Close() {
	for {
		err := p.process.Kill()
		if err != nil {
			log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}
