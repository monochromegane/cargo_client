package command

import (
	"os/exec"
	"strconv"
	"strings"
)

type CargoCommand struct {
	Debug       bool
	Image       string
	User        string
	WorkDir     string
	GroupBy     string
	Mount       string
	Concurrency int
	BeforeAll   string
	Cmd         []string
	Target      string
	Filter      string
}

func (self *CargoCommand) Command() *exec.Cmd {

	cmd := []string{
		"-i", self.Image,
		"-u", self.User,
		"-w", self.WorkDir,
		"-g", self.GroupBy,
		"-m", self.Mount,
		"-n", strconv.Itoa(self.Concurrency),
		"-b", "\"" + self.BeforeAll + "\"",
		"-c", "\"" + strings.Join(self.Cmd, " ") + "\"",
		"-t", self.Target,
		"-f", self.Filter,
	}

	if self.Debug {
		cmd = append(cmd, "--debug")
	}

	return exec.Command("cargo", cmd...)
}
