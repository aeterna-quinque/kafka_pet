package config

type Kafka struct {
	Brokers          []string `env:"BROKERS" envSeparator:","`
	UsersGetTopic    string   `env:"USERS_GET_TOPIC"`
	UsersCreateTopic string   `env:"USERS_CREATE_TOPIC"`
}
