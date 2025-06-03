package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

// Database connection
var db *sql.DB

// URL struct to handle request and response payloads
type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

// Helper function to generate random short codes
func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

// Handler to create a short URL
func createShortURL(w http.ResponseWriter, r *http.Request) {
	var url URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil || url.LongURL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Generate a short code and insert into the database
	shortCode := generateShortCode()
	_, err = db.Exec("INSERT INTO urls (long_url, short_code) VALUES (?, ?)", url.LongURL, shortCode)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	// Return the short URL
	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
	json.NewEncoder(w).Encode(URL{LongURL: url.LongURL, ShortURL: shortURL})
}

// Handler to redirect using a short URL
func redirectToLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["code"]

	var longURL string
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_code = ?", shortCode).Scan(&longURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

// Initialize SQLite database and create table
func initDB() {
	var err error
	// Use "sqlite" instead of "sqlite3"
	db, err = sql.Open("sqlite", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the urls table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        long_url TEXT NOT NULL,
        short_code TEXT NOT NULL UNIQUE
    )`)
	if err != nil {
		log.Fatal(err)
	}
}

// Main Function
func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/shorten", createShortURL).Methods("POST")
	router.HandleFunc("/{code}", redirectToLongURL).Methods("GET")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}