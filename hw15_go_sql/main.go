package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/MoshKillaPit/OtusHomework/hw15_go_sql/db"
)

func isValidURL(endpoint string) bool {
	u, err := url.ParseRequestURI(endpoint)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func createUser(endpoint string, user map[string]string) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}
	postData, _ := json.Marshal(user)
	resp, reqErr := http.Post(endpoint, "application/json", bytes.NewBuffer(postData))
	if reqErr != nil {
		fmt.Println("Ошибка создания пользователя:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", body)
}

func createProduct(endpoint string, product map[string]interface{}) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}
	postData, _ := json.Marshal(product)
	resp, reqErr := http.Post(endpoint, "application/json", bytes.NewBuffer(postData))
	if reqErr != nil {
		fmt.Println("Ошибка создания продукта:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", body)
}

func fetchProducts(endpoint string) {
	if !isValidURL(endpoint) {
		log.Printf("Invalid URL: %s", endpoint)
		return
	}
	resp, reqErr := http.Get(endpoint)
	if reqErr != nil {
		fmt.Println("Ошибка получения списка продуктов:", reqErr)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Products: %s\n", body)
}

func main() {
	// Подключение к базе данных
	dsn := "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable search_path=public"
	database, dbErr := db.NewDB(dsn)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %v", dbErr)
	}
	defer database.Close()

	// Добавление пользователя
	addUserErr := database.AddUser("Alice", "alice@gmail.com", "securepassword")
	if addUserErr != nil {
		log.Printf("Error adding user: %v", addUserErr)
		database.Close()
		return
	}
	fmt.Println("User added successfully!")

	// Получение списка пользователей
	users, getUsersErr := database.GetUsers()
	if getUsersErr != nil {
		log.Printf("Error fetching users: %v", getUsersErr)
		database.Close()
		return
	}
	fmt.Println("Users:", users)

	// Пример добавления продукта
	addProductErr1 := database.AddProduct("Laptop", 1500)
	if addProductErr1 != nil {
		log.Printf("Error adding product: %v", addProductErr1)
		database.Close()
		return
	}
	addProductErr2 := database.AddProduct("Mouse", 50)
	if addProductErr2 != nil {
		log.Printf("Error adding product: %v", addProductErr2)
		database.Close()
		return
	}
	fmt.Println("Product added successfully!")

	// Пример создания заказа с продуктами
	products := []db.OrderProduct{
		{ProductID: 1, Quantity: 1, Price: 1500}, // Laptop
		{ProductID: 2, Quantity: 2, Price: 50},   // Mouse
	}

	placeOrderErr := database.PlaceOrder(1, "2025-01-13", products)
	if placeOrderErr != nil {
		log.Printf("Error placing order: %v", placeOrderErr)
		database.Close()
		return
	}
	fmt.Println("Order placed successfully!")
}
