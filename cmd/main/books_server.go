package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Book represents a book structure
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Sample books data
var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
	{ID: 2, Title: "Clean Code", Author: "Robert Martin", Year: 2008},
	{ID: 3, Title: "Design Patterns", Author: "Gang of Four", Year: 1994},
}

// getBooksHandler handles GET requests to /api/v1/books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Only allow GET method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	
	// Return the books as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Create a new HTTP mux
	mux := http.NewServeMux()
	
	// Register the books endpoint
	mux.HandleFunc("/api/v1/books", getBooksHandler)
	
	// Add a basic health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	// Server configuration
	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}
	
	fmt.Println("Starting HTTPS server on port 443...")
	fmt.Println("Endpoints available:")
	fmt.Println("  GET /api/v1/books - Get all books")
	fmt.Println("  GET /health - Health check")
	
	// Start the server with TLS
	// Note: You'll need valid SSL certificates (server.crt and server.key)
	log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}
