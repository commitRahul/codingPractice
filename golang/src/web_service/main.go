package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Received request")
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
