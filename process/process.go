package process

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"git.resultys.com.br/lib/lower/exception"
)

// Proc estrutura do processo seu estado
type Proc struct {
	process *os.Process
	state   *os.ProcessState
}

// ExecuteAndWait executa o comando e espera sua conclusão independente de sistema operacional
func ExecuteAndWait(command string) *Proc {
	return Command(command, true)
}

// Execute executa o comando e não espera sua conclusão independente de sistema operacional
func Execute(command string) *Proc {
	return Command(command, false)
}

// Command executada um comando no cmd (windows) ou bash(linux) dependendo do ambiente de execução
// Retorna o processo e error
func Command(command string, wait bool) *Proc {
	if runtime.GOOS == "windows" {
		return Cmd(command, wait)
	} else {
		return Bash(command, wait)
	}
}

// Bash executa um comando no linux
// Retorna o processo e error
func Bash(command string, wait bool) *Proc {
	if wait {
		return run("sh", "-c", command)
	}

	return start("sh", "-c", command)
}

// Cmd executa um comando no windows
// Retorna o processo e error
func Cmd(command string, wait bool) *Proc {
	if wait {
		return run("cmd", "/c", command)
	}

	return start("cmd", "/c", command)
}

// RunProgram executa um programa externo e esperar o retorno
func RunProgram(cmd string) *Proc {
	return run(cmd, "", "")
}

// StartProgram executa um programa externo e nao esperar o retorno
func StartProgram(cmd string) *Proc {
	return start(cmd, "", "")
}

func start(cmd, option, parameters string) *Proc {
	process := command(cmd, option, parameters, func(process *exec.Cmd) error {
		return process.Start()
	})

	return process
}

func run(cmd, option, parameters string) *Proc {
	process := command(cmd, option, parameters, func(process *exec.Cmd) error {
		return process.Run()
	})

	return process
}

func command(cmd, option, parameters string, config func(*exec.Cmd) error) *Proc {
	process := exec.Command(cmd, option, parameters)

	err := config(process)
	if err != nil {
		exception.Raise(err.Error(), exception.WARNING)
		return nil
	}

	return &Proc{process: process.Process, state: process.ProcessState}
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
			exception.Raise(err.Error(), exception.WARNING)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}
