package utils

import "net/http"

func CopyHeaders(header http.Header, destination http.ResponseWriter) {
	for headerKey, headerValues := range header {
		for _, headerValue := range headerValues {
			destination.Header().Add(headerKey, headerValue)
		}
	}
}