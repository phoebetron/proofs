package loader

type Config struct {
	Bas string
	Fil string
	Git bool
	Org string
	Rep string
}

func (c Config) Ensure() Config {
	if c.Bas == "" {
		c.Bas = "."
	}

	return c
}

func (c Config) Verify() {
	if c.Fil == "" {
		panic("Config.Fil must not be empty")
	}
	if !c.Git {
		panic("Config.Git must not be empty")
	}
	if c.Org == "" {
		panic("Config.Org must not be empty")
	}
	if c.Rep == "" {
		panic("Config.Rep must not be empty")
	}
}
