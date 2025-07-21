package main

import (
	"fmt"
	//"encoding/json"
	//"log"
	//"net/http"
)

var GSTRate float64 = 0.06
var totalAmount float64 = 100
var GSTAmount float64 = 0
var subTotal float64 = 0

func calcGST(GSTRate float64, totalAmount float64) float64 {
	return totalAmount * GSTRate
}

func main() {
	GSTAmount = calcGST(GSTRate, totalAmount)
	subTotal = totalAmount - GSTAmount

	fmt.Println("Total Amount:", totalAmount)
	fmt.Println("GST Amount:", GSTAmount)
	fmt.Println("Sub Total:", subTotal)
}