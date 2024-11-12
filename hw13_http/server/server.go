package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Book struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var (
	books = []Book{}
	mu    sync.Mutex
)

// createBook обрабатывает POST-запрос для создания книги.
func createBook(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	book := Book{}
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	mu.Lock()
	books = append(books, book)
	mu.Unlock()

	fmt.Printf("Created book %+v\n", book)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
	fmt.Println("Method createBook executed.")
}

// getBooks обрабатывает GET-запрос для получения списка книг.
func getBooks(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	err := json.NewEncoder(w).Encode(books)
	mu.Unlock()
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
	fmt.Println("Method getBooks executed.")
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.Int("port", 8080, "Port number")
	flag.Parse()

	fmt.Println("Server started")
	serverAddress := *ip + ":" + strconv.Itoa(*port)

	http.HandleFunc("/getBooks", getBooks)
	http.HandleFunc("/createBook", createBook)

	// Используем http.Server для добавления таймаутов
	srv := &http.Server{
		Addr:         serverAddress,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
