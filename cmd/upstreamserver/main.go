package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/logger"
)


func main() {
	http.HandleFunc("/test", func(responseWriter http.ResponseWriter, request *http.Request) {
		logger.Info.Printf("📥 Received %s request at %s", request.Method, request.URL.Path)

		for name, values := range request.Header {
			for _, value := range values {
				logger.Info.Printf("🔹 Header: %s = %s", name, value)
			}
		}

		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Error.Printf("❌ Failed to read body: %s", err.Error())
			http.Error(responseWriter, "Failed to read body", http.StatusInternalServerError)
			return
		}
		logger.Info.Printf("📦 Body: %s", string(bodyBytes))

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(map[string]string{
			"status":  "received",
			"message": "Mock upstream response",
		})
	})

	port := "8081"
	log.Printf("🚀 Mock upstream listening on http://localhost:%s/test", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
