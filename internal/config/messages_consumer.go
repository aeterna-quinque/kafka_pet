package config

type MessagesConsumer struct {
	Topics []string `env:"TOPICS,notEmpty" envSeparator:","`
}
