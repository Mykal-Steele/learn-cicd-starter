package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{"Authorization": []string{"ApiKey testKey"}}
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if key != "testKey" {
		t.Errorf("expected 'testKey', got %v", key)
	}

	// this is a test function to verify that CI is working
}
