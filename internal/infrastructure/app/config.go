package app

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `default:"8080"`
	Log  LogConfig
}

type LogConfig struct {
	Level     string `default:"info"`
	Formatter string `default:"json"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("API", &cfg); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return &cfg, nil
}
