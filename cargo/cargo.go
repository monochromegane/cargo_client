package cargo

import (
	"github.com/monochromegane/cargo_client/cargo/command"
	"github.com/monochromegane/cargo_client/cargo/config"
	"path/filepath"
)

type Cargo struct {
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
        mkdir.Command().Run()

	scp := command.ScpCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		From:   cfg.Cargo_Client.SrcDir,
		Host:   cfg.Docker_Host.Host,
		To:     filepath.Join(topath, "current"),
	}
	scp.Command().Run()
}

func (self *Cargo) Run() (result []byte, err error) {
	cfg := self.Config

	cargo := command.CargoCommand{
		Image:       cfg.Docker_Container.Image,
		User:        cfg.Cargo.User,
		WorkDir:     cfg.Cargo.WorkDir,
		GroupBy:     cfg.Cargo.GroupBy,
		Mount:       cfg.Docker_Container.Mount,
		Concurrency: cfg.Cargo.Concurrency,
		Cmd:         cfg.Docker_Container.Command,
	}
	if len(cfg.Go_Package.Package) > 0 {
		cargo.GoPackage = cfg.Go_Package.Package
	}

	ssh := command.SshCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		Host:   cfg.Docker_Host.Host,
		Cmd:    cargo.Command().Args,
	}

	return ssh.Command().Output()
}
