package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResetStockHandler(t *testing.T) {
	CurrentStock = &Stock{}

	var handler = http.HandlerFunc(ResetStockHandler)
	var rr = httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/reset-stock", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
}
