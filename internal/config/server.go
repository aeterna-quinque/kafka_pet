package config

import "time"

type Server struct {
	Host            string        `env:"HOST"`
	Port            uint16        `env:"PORT,notEmpty"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT,notEmpty"`
}
