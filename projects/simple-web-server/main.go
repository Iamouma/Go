package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// structure to hold the message
type Response struct {
	Message string `json:"message"`
}

// Handler for the  /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}

	// set content-type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Route /hello to hellohandler
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}