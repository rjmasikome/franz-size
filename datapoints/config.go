package kafka

type Config struct {
	// General
	TopicPersistence string `koanf:"topicPersistence"`
}

func (c *Config) SetDefaults() {
	c.TopicPersistence = "franz-metrics"
}

func (c *Config) Validate() error {
	return nil
}
