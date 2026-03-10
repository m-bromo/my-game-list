package config

type Config struct {
	Environment string `env:"ENV"`
	PostgresDB  PostgresDB
	Api         API
}

type PostgresDB struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type API struct {
	Host string `env:"API_HOST"`
	Port string `env:"API_PORT"`
}
