package main

import (
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	"github.com/ronfelsenfeld/go-proxy/internal/server"
)


func main() {
	configuration, err := config.Load()
	if err != nil {
		logger.Error.Fatalf("❌ Failed to load configuration: %v", err)
	}

	address := ":" + configuration.Port
	logger.Info.Println("🔒 Starting HTTPS server on port", configuration.Port)

	err = http.ListenAndServeTLS(address, configuration.CertPath, configuration.KeyPath, server.Router())
	if err != nil {
		logger.Error.Fatalf("❌ Failed to start server: %v", err)
	}
}
