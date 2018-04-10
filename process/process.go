package process

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"time"

	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
)

// Proc estrutura do processo seu estado
type Proc struct {
	process *os.Process
	state   *os.ProcessState
}

// Command executada um comando no cmd (windows) ou bash(linux) dependendo do ambiente de execução
// Retorna o processo e error
func Command(command string) (*Proc, error) {
	if runtime.GOOS == "windows" {
		return Cmd(command)
	} else {
		return Bash(command)
	}
}

// Bash executa um comando no linux
// Retorna o processo e error
func Bash(command string) (*Proc, error) {
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

// Cmd executa um comando no windows
// Retorna o processo e error
func Cmd(command string) (*Proc, error) {
	cmd := exec.Command("cmd", "/c", command)
	err := cmd.Run()
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

// Start executa um comando de sistema
// Retorna o processo e o error
func Start(program string) (proc *Proc, err error) {
	cmd := exec.Command(program)
	err1 := cmd.Start()
	if err1 != nil {
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return nil, errors.New("processo nao encontrado")
	}

	return &Proc{process: cmd.Process, state: cmd.ProcessState}, nil
}

// GetID retorna o pid do processo
func (p *Proc) GetID() int {
	return p.process.Pid
}

// Close encerra um processo
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
