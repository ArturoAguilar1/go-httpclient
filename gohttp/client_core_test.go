package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization  Arrange
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type","application/json")
	commonHeaders.Set("User-Agent","cool-http-client")
	client.Headers = commonHeaders

	// Execution  Act
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation Assert

	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}
}