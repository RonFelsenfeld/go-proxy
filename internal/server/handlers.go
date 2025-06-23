package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
)

func pingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	response := map[string]string{"message": "pong"}
	responseWriter.Header().Set("Content-Type", "application/json")

	jsonEncoder := json.NewEncoder(responseWriter)
	jsonEncoder.Encode(response)
}

func proxyHandler(responseWriter http.ResponseWriter, request *http.Request, configuration *config.Config) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		errorMessage := "Invalid request body"
		logger.Error.Printf("❌ %s", errorMessage)
		http.Error(responseWriter, errorMessage, http.StatusBadRequest)
		return
	}

	requestBody[configuration.InjectKey] = configuration.InjectValue
	logger.Info.Printf("🔨 Modified request body: %+v", requestBody)

	modifiedBody, err := json.Marshal(requestBody)
	if err != nil {
		errorMessage := "Failed to encode request body"
		logger.Error.Printf("❌ %s", errorMessage)
		http.Error(responseWriter, errorMessage, http.StatusInternalServerError)
		return
	}

	upstreamRequest, err := http.NewRequest(request.Method, configuration.UpstreamURL, bytes.NewBuffer(modifiedBody))
	if err != nil {
		errorMessage := "Failed to build upstream request"
		logger.Error.Printf("❌ %s", errorMessage)
		http.Error(responseWriter, errorMessage, http.StatusInternalServerError)
		return
	}

	upstreamRequest.Header = request.Header.Clone()

	upstreamResponse, err := http.DefaultClient.Do(upstreamRequest)
	if err != nil {
		errorMessage := "Failed to contact upstream"
		logger.Error.Printf("❌ %s", errorMessage)
		http.Error(responseWriter, errorMessage, http.StatusInternalServerError)
		return
	}
	defer upstreamResponse.Body.Close()

	for headerKey, headerValues := range upstreamResponse.Header {
		for _, headerValue := range headerValues {
			responseWriter.Header().Add(headerKey, headerValue)
		}
	}

	responseWriter.WriteHeader(upstreamResponse.StatusCode)
	io.Copy(responseWriter, upstreamResponse.Body)
}
