package config

type Postgres struct {
	User     string `env:"USER,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
	Host     string `env:"HOST,notEmpty"`
	Port     uint16 `env:"PORT,notEmpty"`
	Name     string `env:"NAME,notEmpty"`
	SSLMode  string `env:"SSLMODE,notEmpty"`
}
