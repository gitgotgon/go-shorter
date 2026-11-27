// Package config ...
package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HTTPPort int
}

// Load configurations
func Load() (Config, error) {
	port := 8080 // default port for local development

	if value := os.Getenv("PORT"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return Config{}, fmt.Errorf("invalid PORT: %w", err)
		}
		port = parsed
	}

	if port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("PORT must be between 1 and 65535")
	}

	return Config{
		HTTPPort: port,
	}, nil
}
