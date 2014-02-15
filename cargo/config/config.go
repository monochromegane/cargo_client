package config

import (
	"io/ioutil"
)

type Config struct {
	Docker_Host struct {
		Ssh_Config string
		Host       string
	}
	Docker_Container struct {
		Image   string
		Mount   string
		Command string
	}
	Cargo struct {
		GroupBy     string
		Concurrency int
		User        string
		WorkDir     string
	}
	Cargo_Client struct {
		SrcDir string
	}
	Go_Package struct {
		Package string
	}
}

func DefaultConfig() *Config {
	cfg := &Config{}

	cfg.Docker_Host.Ssh_Config = "~/.ssh/config"
	cfg.Docker_Host.Host = "default"
	cfg.Cargo.GroupBy = "file-size"
	cfg.Cargo.Concurrency = 1
	cfg.Cargo.User = "cargo"
	cfg.Cargo.WorkDir = "/tmp/cargo"

	cfg.Cargo_Client.SrcDir = "."

	return cfg
}

func CopyFromTemplate() error {
	return ioutil.WriteFile("Cargofile_", []byte(template()), 0644)
}

func template() string {
	return `; Cargofile

[docker-host]
; ssh-config  = ~/.ssh/config
; host        = default

[docker-container]
image       = ; container image
mount       = ; mount volume
command     = ; docker run command

[cargo]
; groupby     = file-size
; user        = cargo
; workdir     = /tmp/cargo
concurrency = 2 ; number of concurrency

[cargo-client]
; srcdir      = .

`
}
