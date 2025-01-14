package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/fixme_my_friend/hw16_docker/repository"
)

// isValidURL проверяет базовую корректность URL,
// возвращает true, если URL подходит.
func isValidURL(endpoint string) bool {
	u, err := url.ParseRequestURI(endpoint)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}

// createUser отправляет POST-запрос в некое API для создания пользователя.
func createUser(endpoint string, user map[string]string) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}
	postData, _ := json.Marshal(user)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(postData))
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		log.Printf("Ошибка создания пользователя: %v", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("[createUser] Response: %s\n", body)
}

// createProduct отправляет POST-запрос в некое API для создания продукта.
func createProduct(endpoint string, product map[string]interface{}) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}
	postData, _ := json.Marshal(product)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(postData))
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		log.Printf("Ошибка создания продукта: %v", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("[createProduct] Response: %s\n", body)
}

// fetchProducts отправляет GET-запрос для получения списка продуктов.
func fetchProducts(endpoint string) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return
	}

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		log.Printf("Ошибка получения списка продуктов: %v", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("[fetchProducts] Products: %s\n", body)
}

func main() {
	// Подключение к базе данных
	dsn := "host=db port=5432 user=postgres password=root dbname=postgres sslmode=disable search_path=public"
	database, dbErr := repository.NewDB(dsn)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %v", dbErr)
	}
	defer database.Close()

	// Настройка сервера API
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      http.NewServeMux(),
	}

	server.Handler.(*http.ServeMux).HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "User endpoint reached!")
	})

	server.Handler.(*http.ServeMux).HandleFunc("/products", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Product endpoint reached!")
	})

	go func() {
		log.Println("Server running on :8080")
		log.Fatal(server.ListenAndServe())
	}()

	// Ожидание запуска сервера
	time.Sleep(2 * time.Second)

	// Выполнение HTTP-запросов
	createUser("http://localhost:8080/users", map[string]string{
		"name":     "AliceFromHTTP",
		"email":    "alice_http@example.com",
		"password": "somepass",
	})

	createProduct("http://localhost:8080/products", map[string]interface{}{
		"name":  "HTTP Product",
		"price": 123,
	})

	fetchProducts("http://localhost:8080/products")
}
