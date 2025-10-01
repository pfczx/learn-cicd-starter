package main

import (
	"testing"
	"net/http"
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)
func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	apiKey, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "my-secret-key" {
		t.Errorf("Expected 'my-secret-key', got '%s'", apiKey)
	}
	t.Errorf("err")
}
