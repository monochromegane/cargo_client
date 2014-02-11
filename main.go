package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/monochromegane/cargo_client/cargo"
	"github.com/monochromegane/cargo_client/cargo/config"
	"os"
)

func main() {
	var cfg config.Config
	err := gcfg.ReadFileInto(&cfg, "Cargofile")
	if err != nil {
		fmt.Printf("Cargofile not found.\n")
		os.Exit(1)
	}
	cargo.Run(cfg)
}
