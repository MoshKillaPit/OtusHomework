package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

type DB struct {
	Conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	// Открытие подключения
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверка соединения
	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Connected to the database successfully")
	return &DB{Conn: conn}, nil
}

func (db *DB) Close() error {
	if db.Conn != nil {
		return db.Conn.Close()
	}
	return nil
}

func (db *DB) AddUser(name, email, password string) error {
	query := `
		INSERT INTO public.users (name, email, password) 
		VALUES ($1, $2, $3)
	`
	_, err := db.Conn.Exec(query, name, email, password)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	fmt.Println("User added:", name)
	return nil
}

func (db *DB) GetUsers() ([]User, error) {
	query := `
		SELECT id, name, email, password 
		FROM users 
	`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if scanErr := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); scanErr != nil {
			return nil, fmt.Errorf("failed to scan user: %w", scanErr)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	fmt.Printf("Fetched %d users\n", len(users))
	return users, nil
}

func (db *DB) AddProduct(name string, price int) error {
	query := `
		INSERT INTO products (name, price) 
		VALUES ($1, $2)
	`
	_, err := db.Conn.Exec(query, name, price)
	if err != nil {
		return fmt.Errorf("failed to add product: %w", err)
	}
	fmt.Println("Product added:", name)
	return nil
}

type Product struct {
	ID    int
	Name  string
	Price int
}

func (db *DB) GetProducts() ([]Product, error) {
	query := `
		SELECT id, name, price 
		FROM products
	`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if scanErr := rows.Scan(&product.ID, &product.Name, &product.Price); scanErr != nil {
			return nil, fmt.Errorf("failed to scan product: %w", scanErr)
		}
		products = append(products, product)
	}

	fmt.Printf("Fetched %d products\n", len(products))
	return products, nil
}

func (db *DB) DeleteProduct(id int) error {
	query := `
		DELETE FROM products 
		WHERE id = $1
	`
	_, err := db.Conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	fmt.Println("Product deleted with ID:", id)
	return nil
}

func (db *DB) AddOrder(userID int, orderDate string, totalAmount int) error {
	query := `
		INSERT INTO orders (user_id, order_date, total_amount) 
		VALUES ($1, $2, $3)
	`
	_, err := db.Conn.Exec(query, userID, orderDate, totalAmount)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	fmt.Printf("Order created for user ID: %d\n", userID)
	return nil
}

type Order struct {
	ID          int
	UserID      int
	OrderDate   string
	TotalAmount int
}

func (db *DB) GetOrdersByUser(userID int) ([]Order, error) {
	query := `
		SELECT id, user_id, order_date, total_amount 
		FROM orders 
		WHERE user_id = $1
	`
	rows, err := db.Conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if scanErr := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount); scanErr != nil {
			return nil, fmt.Errorf("failed to scan order: %w", scanErr)
		}
		orders = append(orders, order)
	}

	fmt.Printf("Fetched %d orders for user ID: %d\n", len(orders), userID)
	return orders, nil
}

func (db *DB) DeleteOrder(id int) error {
	query := `
		DELETE FROM orders 
		WHERE id = $1
	`
	_, err := db.Conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	fmt.Println("Order deleted with ID:", id)
	return nil
}

func (db *DB) PlaceOrder(userID int, orderDate string, products []OrderProduct) error {
	// Начинаем транзакцию
	tx, err := db.Conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Откат транзакции в случае ошибки
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Возобновляем панику
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Добавляем заказ
	var orderID int
	orderQuery := `
		INSERT INTO orders (user_id, order_date, total_amount) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	totalAmount := 0
	for _, product := range products {
		totalAmount += product.Quantity * product.Price
	}
	err = tx.QueryRow(orderQuery, userID, orderDate, totalAmount).Scan(&orderID)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	// Добавляем продукты в заказ
	orderProductQuery := `
		INSERT INTO order_products (order_id, product_id, quantity) 
		VALUES ($1, $2, $3)
	`
	for _, product := range products {
		_, err = tx.Exec(orderProductQuery, orderID, product.ProductID, product.Quantity)
		if err != nil {
			return fmt.Errorf("failed to insert order product: %w", err)
		}
	}

	fmt.Printf("Order %d created successfully for user %d\n", orderID, userID)
	return nil
}

type OrderProduct struct {
	ProductID int
	Quantity  int
	Price     int // Цена продукта для расчёта общей суммы
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
