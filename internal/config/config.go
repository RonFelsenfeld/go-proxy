package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
		log.Println("⚠️  No .env file found. Falling back to system env vars.")
	}

	configuration := &Config{
		Port: 					getEnv("GATEWAY_PORT", "443"),
		CertPath: 			getEnv("GATEWAY_CERT_PATH", "certs/cert.pem"),
		KeyPath: 				getEnv("GATEWAY_KEY_PATH", "certs/key.pem"),
		UpstreamURL: 		getEnv("UPSTREAM_URL", "http://localhost:8081"),
		InjectKey: 			getEnv("INJECT_KEY", "injected_key"),
		InjectValue: 		getEnv("INJECT_VALUE", "injected_value"),
	}

	if err := validateConfiguration(configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}