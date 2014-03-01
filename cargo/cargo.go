package cargo

import (
	"fmt"
	"github.com/monochromegane/cargo_client/cargo/command"
	"github.com/monochromegane/cargo_client/cargo/config"
	"github.com/monochromegane/cargo_client/cargo/option"
	"path/filepath"
	"strings"
)

type Cargo struct {
	Option *option.Option
	Config *config.Config
}

func (self *Cargo) SendAssets() bool {
	cfg := self.Config

	topath := filepath.Join(cfg.Cargo.WorkDir, cfg.Docker_Container.Image, cfg.Cargo.User, "current")
	mkdir := command.SshCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		Host:   cfg.Docker_Host.Host,
		Cmd:    []string{"mkdir", "-p", topath},
	}
	self.printDebug(mkdir.Command().Args)
	result, err := mkdir.Command().CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n%s\n", result, err)
		return false
	}

	git := command.GitLsFilesCommand{}
	tar_c := command.CreateTarCommand{}
	tar_x := command.ExtractTarCommand{topath}
	ssh := command.SshCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		Host:   cfg.Docker_Host.Host,
		Cmd:    tar_x.Command().Args,
	}

	self.printDebug(git.Command().Args)
	self.printDebug(tar_c.Command().Args)
	self.printDebug(ssh.Command().Args)

	stdout, stderr, err := command.Pipeline(
		git.Command(),
		tar_c.Command(),
		ssh.Command(),
	)

	if err != nil {
		fmt.Printf("%s\n%s\n%s\n", stdout, stderr, err)
		return false
	}
	return true
}

func (self *Cargo) Run() (result []byte, err error) {
	cfg := self.Config
	cmd := strings.Split(cfg.Docker_Container.Command, " ")

	cargo := command.CargoCommand{
		Debug:       cfg.Cargo.Debug,
		Image:       cfg.Docker_Container.Image,
		User:        cfg.Cargo.User,
		WorkDir:     cfg.Cargo.WorkDir,
		GroupBy:     cfg.Cargo.GroupBy,
		Mount:       cfg.Docker_Container.Mount,
		Concurrency: cfg.Cargo.Concurrency,
		BeforeAll:   cfg.Docker_Container.Before_All,
		Cmd:         cmd[:len(cmd)-1],
		Target:      cmd[len(cmd)-1],
		Filter:      cfg.Docker_Container.Filter,
	}

	ssh := command.SshCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		Host:   cfg.Docker_Host.Host,
		Cmd:    cargo.Command().Args,
	}

	self.printDebug(ssh.Command().Args)
	return ssh.Command().CombinedOutput()
}

func (self *Cargo) printDebug(log ...interface{}) {
	if self.Option.Debug {
		fmt.Printf("DEBUG: %s\n", log)
	}
}
