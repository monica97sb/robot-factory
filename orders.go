package main

import (
	"io"
	"net/http"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		response, status := Orders(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(response))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func Orders(body io.ReadCloser) (string, int) {
	return `{"order_id": "some-id", "total": 160.11 }`, 201
}
