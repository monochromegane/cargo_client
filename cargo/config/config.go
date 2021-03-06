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
		Image      string
		Mount      string
		Before_All string
		Command    string
		Filter     string
	}
	Cargo struct {
		Debug       bool
		GroupBy     string
		Concurrency int
		User        string
		WorkDir     string
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

	return cfg
}

func CopyFromTemplate() error {
	return ioutil.WriteFile("Cargofile", []byte(template()), 0644)
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

`
}
