package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status: "UP",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to generate response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Application starting on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
