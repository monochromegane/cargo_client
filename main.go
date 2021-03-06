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
	case "init":
		cargoInit()
	case "run":
		run(opts)
	}

}

func run(opts option.Option) {
	cfg := config.DefaultConfig()
	err := gcfg.ReadFileInto(cfg, "Cargofile")
	if err != nil {
		fmt.Printf("Cargofile parse error.\n %s\n", err)
		os.Exit(1)
	}

	cargo := cargo.Cargo{&opts, cfg}
	resultSendAssets := cargo.SendAssets()
	if !resultSendAssets {
		fmt.Println("Failed to send assets to docker host.")
		return
	}

	resultRun, err := cargo.Run()
	fmt.Printf("%s\n", resultRun)
	if err != nil {
		fmt.Printf("Failed to run cargo.\n%s\n", err)
	}
}

func cargoInit() {
	err := config.CopyFromTemplate()
	if err != nil {
		fmt.Printf("Failed to create Cargofile.\n %s\n", err)
		os.Exit(1)
	}
}
