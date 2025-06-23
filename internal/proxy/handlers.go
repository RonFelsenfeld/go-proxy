package proxy

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	apiUtils "github.com/ronfelsenfeld/go-proxy/internal/utils/apiutils"
)

func pingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	response := map[string]string{"message": "pong"}
	responseWriter.Header().Set("Content-Type", "application/json")

	jsonEncoder := json.NewEncoder(responseWriter)
	jsonEncoder.Encode(response)
}

func proxyHandler(responseWriter http.ResponseWriter, request *http.Request, configuration *config.Config) {
	requestBody, err := apiUtils.DecodeJSONBody(request)
	if err != nil {
		logger.Error.Printf("❌ %s", err.Error())
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	requestBody[configuration.InjectKey] = configuration.InjectValue
	logger.Info.Printf("🔨 Modified request body: %+v", requestBody)

	modifiedBody, err := apiUtils.EncodeJSONBody(requestBody)
	if err != nil {
		logger.Error.Printf("❌ %s", err.Error())
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	upstreamResponse, err := apiUtils.ForwardRequest(request, modifiedBody, configuration)
	if err != nil {
		logger.Error.Printf("❌ %s", err.Error())
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer upstreamResponse.Body.Close()

	apiUtils.CopyHeaders(upstreamResponse.Header, responseWriter)

	responseWriter.WriteHeader(upstreamResponse.StatusCode)
	io.Copy(responseWriter, upstreamResponse.Body)
}
