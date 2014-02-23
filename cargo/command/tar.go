package command

import (
	"os/exec"
)

type CreateTarCommand struct {
}

func (self *CreateTarCommand) Command() *exec.Cmd {
	return exec.Command("tar", []string{"-c", "-T", "-", "--null"}...)
}
