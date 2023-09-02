package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port      string `yaml:"port"`
		Debug     bool   `yaml:"debug"`
		LogOutput string `yaml:"log_output"`
		BaseURL   string `yaml:"base_url" env:"BASE_URL" env-required:"true""`
	} `yaml:"app"`

	Db struct {
		DSN string `env-required:"true" env:"POSTGRES_DSN"`
	}
}

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}
