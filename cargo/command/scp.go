package command

import (
	"os/exec"
)

type ScpCommand struct {
	Config string
	From   string
	To     string
	Host   string
}

func (self *ScpCommand) Command() *exec.Cmd {

	cmd := []string{"-r"}

	if len(self.Config) > 0 {
		cmd = append(cmd, []string{"-F", self.Config}...)
	}

	cmd = append(cmd, []string{
		self.From,
		self.Host + ":" + self.To,
	}...)

	return exec.Command("scp", cmd...)
}
