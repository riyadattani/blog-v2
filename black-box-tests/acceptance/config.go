//go:build acceptance

package acceptance

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type TestingConfig struct {
	BaseURL string `envconfig:"BASE_URL" required:"true" default:"http://localhost:8080"`
}

func LoadTestingConfig() (TestingConfig, error) {
	var config TestingConfig

	if err := envconfig.Process("", &config); err != nil {
		return TestingConfig{}, fmt.Errorf("failed to load config - %w", err)
	}

	return config, nil
}
