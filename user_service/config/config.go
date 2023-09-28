package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Port      string `yaml:"port"`
		Debug     bool   `yaml:"debug"`
		LogOutput string `yaml:"logOutput"`
	} `yaml:"app"`

	DB struct {
		DSN string `env:"POSTGRES_DSN" env-required:"true"`
	}
}

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}
