package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
)

type Config struct {
	UpstreamURL string
	UpstreamPort string
	ProxyPort string
	TLSCertPath string
	TLSKeyPath string
	InjectKey string
	InjectValue string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Warn.Println("⚠️  No .env file found. Falling back to system env vars.")
	}

	configuration := &Config{
		UpstreamURL: 					getEnvVar("UPSTREAM_URL", "http://localhost:8081"),
		UpstreamPort: 				getEnvVar("UPSTREAM_PORT", "8081"),
		ProxyPort: 						getEnvVar("PROXY_PORT", "443"),
		TLSCertPath: 					getEnvVar("TLS_CERT_PATH", "certs/cert.pem"),
		TLSKeyPath: 					getEnvVar("TLS_KEY_PATH", "certs/key.pem"),
		InjectKey: 						getEnvVar("INJECT_KEY", "injected_key"),
		InjectValue: 					getEnvVar("INJECT_VALUE", "injected_value"),
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