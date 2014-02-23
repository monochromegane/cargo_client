package command

import (
	"os/exec"
)

type pipe struct {
	commands []*exec.Cmd
}

func NewPipe() *pipe {
	return &pipe{}
}

func (self *pipe) Pipe(command *exec.Cmd) *pipe {
	self.commands = append(self.commands, command)
	return self
}

func (self *pipe) Command() *exec.Cmd {
	var args []string
	var path string
	for i, cmd := range self.commands {
		if i == 0 {
			path = cmd.Path
		} else {
			args = append(args, cmd.Path)
		}
		args = append(args, cmd.Args[1:]...)
		if i != len(self.commands)-1 {
			args = append(args, "|")
		}
	}
	return exec.Command(path, args...)
}

// pipe(git).pipe(tar).pipe(ssh).Command().Run()
