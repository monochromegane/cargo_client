package cargo

import (
	"github.com/monochromegane/cargo_client/cargo/command"
	"github.com/monochromegane/cargo_client/cargo/config"
	"path/filepath"
)

type Cargo struct {
	Config config.Config
}

func (self *Cargo) SendAssets() {
	cfg := self.Config
	scp := command.ScpCommand{
		Config: cfg.Docker_Host.Ssh_Config,
		From:   cfg.Cargo_Client.SrcDir,
		Host:   cfg.Docker_Host.Host,
		To:     filepath.Join(cfg.Cargo.WorkDir, cfg.Docker_Container.Image, cfg.Cargo.User, "current"),
	}
        scp.Command().Run()
}
