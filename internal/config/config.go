package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
)

type Config struct {
	Port string
	CertPath string
	KeyPath string
	UpstreamURL string
	InjectKey string
	InjectValue string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Warn.Println("⚠️  No .env file found. Falling back to system env vars.")
	}

	configuration := &Config{
		Port: 					getEnvVar("GATEWAY_PORT", "443"),
		CertPath: 			getEnvVar("GATEWAY_CERT_PATH", "certs/cert.pem"),
		KeyPath: 				getEnvVar("GATEWAY_KEY_PATH", "certs/key.pem"),
		UpstreamURL: 		getEnvVar("UPSTREAM_URL", "http://localhost:8081"),
		InjectKey: 			getEnvVar("INJECT_KEY", "injected_key"),
		InjectValue: 		getEnvVar("INJECT_VALUE", "injected_value"),
	}

	if err := validateConfiguration(configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func getEnvVar(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}