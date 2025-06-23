package upstreamserver

import (
	"encoding/json"
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	apiUtils "github.com/ronfelsenfeld/go-proxy/internal/utils/apiutils"
)

func testHandler(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Info.Printf("📥 Received %s request at %s", request.Method, request.URL.Path)

	apiUtils.PrintRequestHeaders(request)

	requestBody, err := apiUtils.DecodeRequestBody(request)
	if err != nil {
		logger.Error.Printf("❌ %s", err.Error())
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	logger.Info.Printf("📦 Body: %v", requestBody)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	json.NewEncoder(responseWriter).Encode(map[string]any{
		"status":  "success",
		"body": requestBody,
	})
}