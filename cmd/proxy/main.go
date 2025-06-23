package main

import (
	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
)


func main() {
	configuration, err := config.Load()

	if err != nil {
		logger.Error.Fatalf("❌ Failed to load configuration: %v", err)
	}

	logger.Info.Printf("✅ Loaded configuration: %+v\n", configuration)
}
