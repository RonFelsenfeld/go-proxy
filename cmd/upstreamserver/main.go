package main

import (
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	"github.com/ronfelsenfeld/go-proxy/internal/upstreamserver"
)

func main() {
	configuration, err := config.Load()
	if err != nil {
		logger.Error.Fatalf("❌ Failed to load configuration: %v", err)
	}

	address := ":" + configuration.UpstreamPort
	
	logger.Info.Printf("🚀 Mock upstream server listening on http://localhost:%s/test", configuration.UpstreamPort)
	err = http.ListenAndServe(address, upstreamserver.Router())

	if err != nil {
		logger.Error.Fatalf("❌ Failed to start upstream server: %v", err)
	}
}
