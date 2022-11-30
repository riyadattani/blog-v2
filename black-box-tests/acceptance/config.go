//go:build acceptance

package acceptance

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type testingConfig struct {
	BaseURL string `envconfig:"BASE_URL" required:"true" default:"http://localhost:8080"`
}

func loadTestingConfig() (testingConfig, error) {
	var config testingConfig

	if err := envconfig.Process("", &config); err != nil {
		return testingConfig{}, fmt.Errorf("failed to load config - %w", err)
	}

	return config, nil
}
