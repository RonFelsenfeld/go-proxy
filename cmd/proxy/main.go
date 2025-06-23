package main

import (
	"log"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
)


func main() {
	configuration, err := config.Load()

	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	log.Printf("✅ Loaded config: %+v\n", configuration)
}
