package command

import (
	"os/exec"
	"strconv"
)

type CargoCommand struct {
	Image       string
	User        string
	WorkDir     string
	GroupBy     string
	Mount       string
	Concurrency int
	Cmd         string
	GoPackage   string
}

func (self *CargoCommand) Command() *exec.Cmd {

	cmd := []string{
		"-i", self.Image,
		"-u", self.User,
		"-w", self.WorkDir,
		"-g", self.GroupBy,
		"-m", self.Mount,
		"-n", strconv.Itoa(self.Concurrency),
		"-c", "\"" + self.Cmd + "\"",
	}

	if len(self.GoPackage) > 0 {
		cmd = append(cmd, []string{"--go-package", self.GoPackage}...)
	}

	return exec.Command("/vagrant/go/bin/cargo", cmd...)
}
