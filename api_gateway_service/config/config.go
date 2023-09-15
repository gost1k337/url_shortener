package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port      string `yaml:"port"`
		Host      string `yaml:"host"`
		Debug     bool   `yaml:"debug"`
		LogOutput string `yaml:"log_output"`
	} `yaml:"app"`

	UrlShorteningService struct {
		Port string `yaml:"port"`
	} `yaml:"url_shortening_service"`

	UserService struct {
		Port string `yaml:"port"`
	} `yaml:"user_service"`
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
