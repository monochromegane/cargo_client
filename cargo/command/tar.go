package command

import (
	"os/exec"
)

type CreateTarCommand struct {
}

type ExtractTarCommand struct {
	DestDir string
}

func (self *CreateTarCommand) Command() *exec.Cmd {
	return exec.Command("tar", []string{"-c", "-T", "-", "--null"}...)
}

func (self *ExtractTarCommand) Command() *exec.Cmd {
	return exec.Command("tar", []string{"x", "-C", self.DestDir}...)
}
