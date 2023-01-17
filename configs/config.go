package configs

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"net/url"
)

type Config struct {
	ServerAddress string `env:"HTTP_SERVER_ADDRESS" envDefault:"http://0.0.0.0:8080"`
}

func (c *Config) GetServerAddress() (*url.URL, error) {
	return url.Parse(c.ServerAddress)
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	cfn := &Config{}
	if err := env.Parse(cfn); err != nil {
		return nil, err
	}

	return cfn, nil
}
