package main

import (
	"encoding/json"
	"net/http"
)

var CurrentStock *Stock

type Stock struct {
	Components []Component `json:"components"`
}

type Component struct {
	Code      string  `json:"code"`
	Price     float32 `json:"price"`
	Available int     `json:"available"`
	Part      string  `json:"part"`
}

func (m *Stock) ToJSON() string {
	o, err := json.MarshalIndent(&m, "", "\t")
	if err != nil {
		return "Error in conversion"
	}
	return string(o)
}

func ResetStockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		response, status := ResetStock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(response))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func ResetStock() (string, int) {
	CurrentStock = &Stock{}

	codes := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	prices := []float32{10.28, 24.07, 13.30, 28.94, 12.39, 30.77, 55.13, 50.00, 90.12, 82.31}
	availables := []int{9, 7, 0, 1, 3, 2, 15, 7, 92, 15}
	parts := []string{"Humanoid Face", "LCD Face", "Steampunk Face", "Arms with Hands",
		"Arms with Grippers", "Mobility with Wheels", "Mobility with Legs", "Mobility with Tracks",
		"Material Bioplastic", "Material Metallic"}

	for i := range codes {
		component := Component{
			Code:      codes[i],
			Price:     prices[i],
			Available: availables[i],
			Part:      parts[i],
		}
		CurrentStock.Components = append(CurrentStock.Components, component)
	}

	return CurrentStock.ToJSON(), 201
}
