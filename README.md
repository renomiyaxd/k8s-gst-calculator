# GST Calculator API

A simple REST API written in Go that calculates GST (Goods and Services Tax) based on a total amount.

## 📦 How to Run Locally

```bash
cd source
go run main.go
```

## Sample Request
```bash
http://localhost:8080/calculate?amount=100&gst=0.06
```

## Sample Response
```bash
{"total_amount":100,"gst_rate":0.06,"gst_amount":4.00,"sub_total":94.00}
```
