package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string `env:"BOT_TOKEN"`
	DBConfig
}

type DBConfig struct {
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
