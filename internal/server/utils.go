package server

import (
	"net/http"
)

func getIsPostRequest(request *http.Request) bool {
	return request.Method == http.MethodPost
}

func getIsPutRequest(request *http.Request) bool {
	return request.Method == http.MethodPut
}
