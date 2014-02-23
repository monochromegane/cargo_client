package command

import (
	"os/exec"
)

type GitLsFilesCommand struct {
}

func (self *GitLsFilesCommand) Command() *exec.Cmd {
	return exec.Command("git", []string{"ls-files", "-z"}...)
}
