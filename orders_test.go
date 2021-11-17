package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewOrderHandler(t *testing.T) {
	var handler = http.HandlerFunc(OrdersHandler)
	var rr = httptest.NewRecorder()

	body := `{"components": ["I", "A", "D", "F"]}`
	req, err := http.NewRequest("POST", "/orders", bytes.NewReader([]byte(body)))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected := `{"order_id": "some-id", "total": 160.11 }`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
