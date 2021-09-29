package endpoints

type Config struct {
	// General
	Endpoint string `koanf:"endpoint"`
}

func (c *Config) SetDefaults() {
	c.Endpoint = "metrics"
}

func (c *Config) Validate() error {
	return nil
}
