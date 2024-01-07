package app

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `default:"8080"`
	Log  LogConfig
	Grpc GRPCConfig
}

type LogConfig struct {
	Level     string `default:"info"`
	Formatter string `default:"json"`
}

type GRPCConfig struct {
	Port string `default:"5005"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("API", &cfg); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return &cfg, nil
}
