package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
