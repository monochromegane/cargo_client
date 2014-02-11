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
		cfg.Ssh.SshConfig,
		cfg.Docker.Src,
		cfg.Ssh.Host+":"+filepath.Join(cfg.Cargo.WorkDir, cfg.Docker.Image, cfg.Cargo.User, "current"),
	)
        // fmt.Printf("%s\n", cmd.Args)
        cmd.Run()
}

func Ssh(cfg config.Config){
	cmd := exec.Command(
		"ssh",
		"-F",
		cfg.Ssh.SshConfig,
                cfg.Ssh.Host,
                "/vagrant/go/bin/cargo",
                "-i",
                cfg.Docker.Image,
                "-u",
                cfg.Cargo.User,
                "-w",
                cfg.Cargo.WorkDir,
                "-g",
                cfg.Cargo.GroupBy,
                "--go-package",
                cfg.GoPackage.Package,
                "-d",
                cfg.Docker.Dest,
                "-n",
                strconv.Itoa(cfg.Cargo.Concurrency),
                "-c",
                "\"" + cfg.Docker.Command + "\"",
	)
        // fmt.Printf("%s\n", cmd.Args)
        result, err := cmd.Output()
        fmt.Printf("%s, %s\n", result, err)
}
