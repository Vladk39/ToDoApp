package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBconfig   string `env:"DBconfig"`
	ServerPort string `env:"SERVERPORT"`
}

func GetConfig() (*Config, error) {
	err := godotenv.Load("../envs/.env")
	if err != nil {
		fmt.Println("Ошибка загрузки .env файла")
	}

	conf := &Config{}

	err = env.Parse(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
