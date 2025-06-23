package utils

import "net/http"

func CopyHeaders(source http.Header, destination http.ResponseWriter) {
	for headerKey, headerValues := range source {
		for _, headerValue := range headerValues {
			destination.Header().Add(headerKey, headerValue)
		}
	}
}