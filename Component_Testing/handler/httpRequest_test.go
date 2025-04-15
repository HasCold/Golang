package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// testing.T --> Normal Testing
// testing.B --> Bench Mark Testing
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder() // Prepare the writer
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected Status Code %d but got %d", http.StatusOK, status)
	}

	var res Response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World"
	if res.Message != expected {
		t.Errorf("Expected Response %s but got this response %s", expected, res.Message)
	}
}
