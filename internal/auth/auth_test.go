package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Authorization", "ApiKey test_api_key")
	got, error := GetAPIKey(headers)
	want := "test_api_key"
	if error != nil {
		t.Fatalf("expected no error, got %v", error)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected = %v, got %v", want, got)
	}
}
