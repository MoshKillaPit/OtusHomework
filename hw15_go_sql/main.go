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

	"github.com/MoshKillaPit/OtusHomework/hw15_go_sql/repository"
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
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		fmt.Println("Ошибка создания пользователя:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[createUser] Response: %s\n", body)
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
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		fmt.Println("Ошибка создания продукта:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[createProduct] Response: %s\n", body)
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
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		fmt.Println("Ошибка получения списка продуктов:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[fetchProducts] Products: %s\n", body)
}

func main() {
	// Подключение к базе данных
	dsn := "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable search_path=public"
	database, dbErr := repository.NewDB(dsn)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %v", dbErr)
	}
	defer database.Close()

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
