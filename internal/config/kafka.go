package config

type Kafka struct {
	Brokers       []string `env:"BROKERS" envSeparator:","`
	UsersTopic    string   `env:"USERS_TOPIC"`
	RequestsTopic string   `env:"REQUESTS_TOPIC"`
}
