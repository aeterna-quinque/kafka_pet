package config

type Kafka struct {
	Brokers     []string `env:"BROKERS" envSeparator:","`
	ServerTopic string   `env:"SERVER_TOPIC"`
}
