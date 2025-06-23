package upstreamserver

import "net/http"

func Router() http.Handler {
	upstreamServerRouter := http.NewServeMux()

	upstreamServerRouter.HandleFunc("/test", testHandler)

	return upstreamServerRouter
}