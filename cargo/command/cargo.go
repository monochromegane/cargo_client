package command

import (
	"os/exec"
	"strconv"
	"strings"
)

type CargoCommand struct {
	Image       string
	User        string
	WorkDir     string
	GroupBy     string
	Mount       string
	Concurrency int
	Cmd         []string
	Target      string
}

func (self *CargoCommand) Command() *exec.Cmd {

	cmd := []string{
		"-i", self.Image,
		"-u", self.User,
		"-w", self.WorkDir,
		"-g", self.GroupBy,
		"-m", self.Mount,
		"-n", strconv.Itoa(self.Concurrency),
		"-c", "\"" + strings.Join(self.Cmd, " ") + "\"",
		"-t", self.Target,
	}

	return exec.Command("/vagrant/go/bin/cargo", cmd...)
}
