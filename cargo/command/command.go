package command

import (
	"os/exec"
        "fmt"
        "github.com/monochromegane/cargo_client/cargo/config"
        "strconv"
        "path/filepath"
)

func Scp(cfg config.Config) {
	cmd := exec.Command(
		"scp",
                "-r",
		"-F",
		cfg.Docker_Host.Ssh_Config,
		cfg.Cargo_Client.SrcDir,
		cfg.Docker_Host.Host+":"+filepath.Join(cfg.Cargo.WorkDir, cfg.Docker_Container.Image, cfg.Cargo.User, "current"),
	)
        // fmt.Printf("%s\n", cmd.Args)
        cmd.Run()
}

func Ssh(cfg config.Config){
	cmd := exec.Command(
		"ssh",
		"-F",
		cfg.Docker_Host.Ssh_Config,
                cfg.Docker_Host.Host,
                "/vagrant/go/bin/cargo",
                "-i",
                cfg.Docker_Container.Image,
                "-u",
                cfg.Cargo.User,
                "-w",
                cfg.Cargo.WorkDir,
                "-g",
                cfg.Cargo.GroupBy,
                "--go-package",
                cfg.Go_Package.Package,
                "-d",
                cfg.Docker_Container.Mount,
                "-n",
                strconv.Itoa(cfg.Cargo.Concurrency),
                "-c",
                "\"" + cfg.Docker_Container.Command + "\"",
	)
        // fmt.Printf("%s\n", cmd.Args)
        result, err := cmd.Output()
        fmt.Printf("%s, %s\n", result, err)
}
