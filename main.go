package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/monochromegane/cargo_client/cargo"
        "os"
)

func main() {
        var cfg cargo.Config
        err := gcfg.ReadFileInto(&cfg, "Cargofile")
        if err != nil {
                fmt.Printf("Cargofile not found.\n")
                os.Exit(1)
        }
        fmt.Printf("%s\n", cfg)
}
