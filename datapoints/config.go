package datapoints

import "time"

type Config struct {
	// General
	TopicPersistence string        `koanf:"topicPersistence"`
	ProbeInterval    time.Duration `koanf:"probeInterval"`
}

func (c *Config) SetDefaults() {
	c.TopicPersistence = "franz-metrics"
	c.ProbeInterval = 60 * time.Second
}

func (c *Config) Validate() error {
	return nil
}
