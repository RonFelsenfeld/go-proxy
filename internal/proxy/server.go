package proxy

import (
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
	"github.com/ronfelsenfeld/go-proxy/internal/logger"
	apiUtils "github.com/ronfelsenfeld/go-proxy/internal/utils/apiutils"
)

func Router(configuration *config.Config) http.Handler {
	proxyRouter := http.NewServeMux()

	proxyRouter.HandleFunc("/ping", pingHandler)

	proxyRouter.HandleFunc("/proxy", func(responseWriter http.ResponseWriter, request *http.Request) {
		logger.Info.Println("🔍 Request received:", request.Method, request.URL.Path)

	
		if apiUtils.GetIsPostRequest(request) || apiUtils.GetIsPutRequest(request) {
			proxyHandler(responseWriter, request, configuration)
		} else {
			logger.Error.Println("⛔️ Method not allowed:", request.Method, request.URL.Path)
			http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return proxyRouter
}
