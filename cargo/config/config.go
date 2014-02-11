package config

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
