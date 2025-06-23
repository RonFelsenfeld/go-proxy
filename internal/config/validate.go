package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func validateConfiguration(configuration *Config) error {
	if strings.TrimSpace(configuration.UpstreamURL) == "" {
		return errors.New("UPSTREAM_URL is required")
	}

	if err := getIsFileExists(configuration.CertPath); err != nil {
		return fmt.Errorf("invalid TLS_CERT_PATH: %w", err)
	}

	if err := getIsFileExists(configuration.KeyPath); err != nil {
		return fmt.Errorf("invalid TLS_KEY_PATH: %w", err)
	}

	return nil
}

func getIsFileExists(path string) error {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", path)
	}

	if info.IsDir() {
		return fmt.Errorf("path is a directory, not a file: %s", path)
	}

	return nil
}