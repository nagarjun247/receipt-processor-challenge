package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPointsHandler).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
