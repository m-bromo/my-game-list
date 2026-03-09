package config

type Config struct {
	Environment string
	PostgresDB  PostgresDB
}

type PostgresDB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type API struct {
	Host string
	Port string
}
