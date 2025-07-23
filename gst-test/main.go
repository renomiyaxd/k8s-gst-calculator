//renomiya gst calculator 2025

package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
)

type GSTResponse struct {
	TotalAmount float64 `json:"total_amount"`
	GSTRate     float64 `json:"gst_rate"`
	GSTAmount   float64 `json:"gst_amount"`
	SubTotal    float64 `json:"sub_total"`
}

func calcGST(GSTRate float64, totalAmount float64) (gstAmount float64, subTotal float64) {
	gstAmount = math.Round((totalAmount * GSTRate) * 100) / 100
	subTotal = math.Round((totalAmount - gstAmount) * 100) / 100
	return
}

func GSTHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount. Please enter a valid amount in decimal format", http.StatusBadRequest)
		return
	}
	GSTStr := r.URL.Query().Get("gst")
	GST, err := strconv.ParseFloat(GSTStr, 64)
	if err != nil || GST >= 1{
		http.Error(w, "Invalid GST rate. Please enter a valid GST rate between 0 to 0.99 in decimal format", http.StatusBadRequest)
		return
	}

	gstAmount, subTotal := calcGST(GST, amount)

	response := GSTResponse{
		TotalAmount: amount,
		GSTRate: GST,
		GSTAmount: gstAmount,
		SubTotal: subTotal,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	
}

func main() {
	http.HandleFunc("/calculate", GSTHandler)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}