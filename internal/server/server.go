package server

import (
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	"github.com/ronfelsenfeld/go-proxy/internal/utils"
)

func Router(configuration *config.Config) http.Handler {
	requestRouter := http.NewServeMux()

	requestRouter.HandleFunc("/ping", pingHandler)

	requestRouter.HandleFunc("/proxy", func(responseWriter http.ResponseWriter, request *http.Request) {
		logger.Info.Println("🔍 Request received:", request.Method, request.URL.Path)
		
		if utils.GetIsPostRequest(request) || utils.GetIsPutRequest(request) {
			proxyHandler(responseWriter, request, configuration)
		} else {
			logger.Error.Println("⛔️ Method not allowed:", request.Method, request.URL.Path)
			http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return requestRouter
}
