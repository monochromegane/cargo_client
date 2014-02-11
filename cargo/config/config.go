package config

type Config struct {
	Docker struct {
		Src     string
		Dest    string
		Command string
		Image   string
	}
	Cargo struct {
		GroupBy     string
		Concurrency int
		User        string
		WorkDir     string
	}
	Ssh struct {
		SshConfig string
		Host      string
	}
	GoPackage struct {
		Package string
	}
}


