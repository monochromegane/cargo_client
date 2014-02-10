package cargo

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
        GoPackage struct {
                Package string
        }
}
