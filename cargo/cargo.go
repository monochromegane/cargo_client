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

func (self *Cargo) SendAssets() {
	cfg := self.Config

	topath := filepath.Join(cfg.Cargo.WorkDir, cfg.Docker_Container.Image, cfg.Cargo.User)
	mkdir := command.SshCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		Host:   cfg.Docker_Host.Host,
		Cmd:    []string{"mkdir", "-p", topath},
	}
	self.printDebug(mkdir.Command().Args)
	mkdir.Command().Run()

	scp := command.ScpCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		From:   cfg.Cargo_Client.SrcDir,
		Host:   cfg.Docker_Host.Host,
		To:     filepath.Join(topath, "current"),
	}
	self.printDebug(scp.Command().Args)
	scp.Command().Run()
}

func (self *Cargo) Run() (result []byte, err error) {
	cfg := self.Config
	cmd := strings.Split(cfg.Docker_Container.Command, " ")

	cargo := command.CargoCommand{
		Image:       cfg.Docker_Container.Image,
		User:        cfg.Cargo.User,
		WorkDir:     cfg.Cargo.WorkDir,
		GroupBy:     cfg.Cargo.GroupBy,
		Mount:       cfg.Docker_Container.Mount,
		Concurrency: cfg.Cargo.Concurrency,
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
	return ssh.Command().Output()
}

func (self *Cargo) printDebug(log ...interface{}) {
	if self.Option.Debug {
		fmt.Printf("DEBUG: %s\n", log)
	}
}
