package cargo

import (
	"github.com/monochromegane/cargo_client/cargo/command"
	"github.com/monochromegane/cargo_client/cargo/config"
)

type Cargo struct {
	Config config.Config
}

func (self *Cargo) SendAssets() {
	command.Scp(self.Config)
}

func (self *Cargo) Run() {
	command.Ssh(self.Config)
}
