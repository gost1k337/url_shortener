package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Port      string `yaml:"port"`
		Host      string `yaml:"host"`
		Debug     bool   `yaml:"debug"`
		LogOutput string `yaml:"logOutput"`
	} `yaml:"app"`

	URLShorteningService struct {
		Port string `yaml:"port"`
		Host string `yaml:"host" env:"URL_SHORTENING_SERVICE_HOST" env-required:"true"`
	} `yaml:"urlShorteningService"`

	UserService struct {
		Port string `yaml:"port"`
		Host string `yaml:"host" env:"USER_SERVICE_HOST" env-required:"true"`
	} `yaml:"userService"`
}

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}
