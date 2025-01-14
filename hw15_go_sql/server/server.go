package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MoshKillaPit/OtusHomework/hw15_go_sql/db" // Импорт пакета для работы с базой данных
)

var database *db.DB

func main() {
	// Подключение к базе данных
	dsn := "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable search_path=public"
	var err error
	database, err = db.NewDB(dsn)
	if err != nil {
		// Вместо log.Fatalf используем log.Printf + return
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer database.Close()

	// Настройка маршрутов
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/products", handleProducts)
	http.HandleFunc("/orders", handleOrders)

	// Запуск сервера
	serverAddress := "127.0.0.1:8080"
	fmt.Println("Server started on", serverAddress)

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	// В случае ошибки – просто логируем и выходим,
	// что даст сработать defer-позовам (закрыть базу и т.д.)
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
		return
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Добавление пользователя
		var user struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := database.AddUser(user.Name, user.Email, user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User added successfully")
	case http.MethodGet:
		// Получение списка пользователей
		users, err := database.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Добавление продукта
		var product struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := database.AddProduct(product.Name, product.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Product added successfully")
	case http.MethodGet:
		// Получение списка продуктов
		products, err := database.GetProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Создание заказа
		var order struct {
			UserID    int               `json:"userid"`
			OrderDate string            `json:"orderdate"`
			Products  []db.OrderProduct `json:"products"`
		}
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := database.PlaceOrder(order.UserID, order.OrderDate, order.Products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Order placed successfully")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
