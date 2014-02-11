package command

import (
	"os/exec"
)

type SshCommand struct {
	Config string
	Host   string
	Cmd    []string
}

func (self *SshCommand) Command() *exec.Cmd {

	cmd := []string{}

	if len(self.Config) > 0 {
		cmd = append(cmd, []string{"-F", self.Config}...)
	}

	cmd = append(cmd, self.Host)
	cmd = append(cmd, self.Cmd...)

	return exec.Command("ssh", cmd...)
}
