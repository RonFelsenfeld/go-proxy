package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ronfelsenfeld/go-proxy/internal/config"
)

func GetIsPostRequest(request *http.Request) bool {
	return request.Method == http.MethodPost
}

func GetIsPutRequest(request *http.Request) bool {
	return request.Method == http.MethodPut
}

func DecodeJSONBody(request *http.Request) (map[string]any, error) {
	var decodedBody map[string]any

	if err := json.NewDecoder(request.Body).Decode(&decodedBody); err != nil {
		return nil, errors.New("invalid request body")
	}

	return decodedBody, nil
}

func EncodeJSONBody(body map[string]any) ([]byte, error) {
	encodedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.New("failed to encode request body")
	}

	return encodedBody, nil
}

func ForwardRequest(originalRequest *http.Request, requestBody []byte, configuration *config.Config) (*http.Response, error) {
	upstreamRequest, err := http.NewRequest(
		originalRequest.Method,
		configuration.UpstreamURL,
		bytes.NewBuffer(requestBody),
	)

	if err != nil {
		return nil, errors.New("failed to build upstream request")
	}

	upstreamRequest.Header = originalRequest.Header.Clone()

	upstreamResponse, err := http.DefaultClient.Do(upstreamRequest)
	if err != nil {
		return nil, errors.New("failed to contact upstream")
	}

	return upstreamResponse, nil
}

