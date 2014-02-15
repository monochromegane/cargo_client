package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/cargo_client/cargo"
	"github.com/monochromegane/cargo_client/cargo/config"
	"github.com/monochromegane/cargo_client/cargo/option"
	"os"
)

var opts option.Option

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "cargo"
	parser.Usage = "[OPTIONS] run"

	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	switch args[0] {
	case "run":
		run()
	}

}

func run() {
	cfg := config.DefaultConfig()
	err := gcfg.ReadFileInto(cfg, "Cargofile")
	if err != nil {
		fmt.Printf("Cargofile parse error.\n %s\n", err)
		os.Exit(1)
	}
	cargo := cargo.Cargo{cfg}
	cargo.SendAssets()
	result, err := cargo.Run()
	fmt.Printf("%s, %s\n", result, err)
}
