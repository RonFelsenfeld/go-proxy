package server

import (
	"encoding/json"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(responseWriter http.ResponseWriter, request *http.Request) {
		response := map[string]string{"message": "pong"}
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(response)
	})

	return mux
}