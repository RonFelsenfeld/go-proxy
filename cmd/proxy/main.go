package main

import (
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	"github.com/ronfelsenfeld/go-proxy/internal/proxy"
)

func main() {
	configuration, err := config.Load()
	if err != nil {
		logger.Error.Fatalf("❌ Failed to load configuration: %v", err)
	}

	address := ":" + configuration.ProxyPort

	logger.Info.Printf("🚀 Proxy server listening on https://localhost:%s/proxy", configuration.ProxyPort)
	err = http.ListenAndServeTLS(address, configuration.TLSCertPath, configuration.TLSKeyPath, proxy.Router(configuration))

	if err != nil {
		logger.Error.Fatalf("❌ Failed to start proxy server: %v", err)
	}
}
