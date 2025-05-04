package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Create a request with an Authorization header
	header := http.Header{}

	header.Add("Authorization", "ApiKey test-APIKey")

	// Call the function
	apiKey, _ := GetAPIKey(header)

	// Verify the result
	expectedKey := "test-APIKey"
	if apiKey != expectedKey {
		t.Errorf("Expected API key to be %s, got %s", expectedKey, apiKey)
	}
}

func TestGetAPIKeyError(t *testing.T) {
	// Create a request with an Authorization header
	header := http.Header{}

	header.Add("Authorization", "ApiKey")

	// Call the function
	_, err := GetAPIKey(header)

	// Verify the result
	if err == nil {
		t.Error("Expected an error but got nil")
	}
}

func TestGetAPIKeyErrorEmpty(t *testing.T) {
	// Create a request with an Authorization header
	header := http.Header{}

	header.Add("Authorization", "")

	// Call the function
	apiKey, _ := GetAPIKey(header)

	// Verify the result
	if apiKey != "" {
		t.Errorf("Expected empty API key, got %s", apiKey)
	}
}
