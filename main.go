package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/orders", OrdersHandler)
	http.HandleFunc("/reset-stock", ResetStockHandler)
	http.ListenAndServe(":8080", nil)
}
