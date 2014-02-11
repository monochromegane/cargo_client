package cargo

import (
	"github.com/monochromegane/cargo_client/cargo/config"
	"github.com/monochromegane/cargo_client/cargo/command"
)

func Run(cfg config.Config) {
	command.Scp(cfg)
	command.Ssh(cfg)
}
