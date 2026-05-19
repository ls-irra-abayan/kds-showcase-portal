package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandlerReturnsKDSShowcaseContent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	homeHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	body := recorder.Body.String()
	expectedContent := []string{
		"Kitchen Display System (KDS)",
		"Overview",
		"Demos",
		"Key Features",
		"Bugs by Label and Priority",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(body, expected) {
			t.Errorf("expected response body to include %q", expected)
		}
	}
}

func TestHomeHandlerReturnsNotFoundForNonRootPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	homeHandler(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}
