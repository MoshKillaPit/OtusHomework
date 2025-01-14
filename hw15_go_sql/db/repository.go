package db

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	conn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{conn: conn}
}

// User-related queries
func (r *Repository) AddUser(user User) error {
	query := `
		INSERT INTO public.users (name, email, password) 
		VALUES ($1, $2, $3)
	`
	_, err := r.conn.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	return nil
}

func (r *Repository) GetUsers() ([]User, error) {
	query := `
		SELECT id, name, email, password 
		FROM users
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// Product-related queries
func (r *Repository) AddProduct(product Product) error {
	query := `
		INSERT INTO products (name, price) 
		VALUES ($1, $2)
	`
	_, err := r.conn.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("failed to add product: %w", err)
	}
	return nil
}

func (r *Repository) GetProducts() ([]Product, error) {
	query := `
		SELECT id, name, price 
		FROM products
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *Repository) DeleteProduct(id int) error {
	query := `
		DELETE FROM products 
		WHERE id = $1
	`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

// Order-related queries
func (r *Repository) AddOrder(order Order) (int, error) {
	query := `
		INSERT INTO orders (user_id, order_date, total_amount) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	var orderID int
	err := r.conn.QueryRow(query, order.UserID, order.OrderDate, order.TotalAmount).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert order: %w", err)
	}
	return orderID, nil
}

func (r *Repository) GetOrdersByUser(userID int) ([]Order, error) {
	query := `
		SELECT id, user_id, order_date, total_amount 
		FROM orders 
		WHERE user_id = $1
	`
	rows, err := r.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *Repository) DeleteOrder(id int) error {
	query := `
		DELETE FROM orders 
		WHERE id = $1
	`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}

func (r *Repository) AddOrderProduct(orderID int, product OrderProduct) error {
	query := `
		INSERT INTO order_products (order_id, product_id, quantity, price) 
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.conn.Exec(query, orderID, product.ProductID, product.Quantity, product.Price)
	if err != nil {
		return fmt.Errorf("failed to insert order product: %w", err)
	}
	return nil
}

type Product struct {
	ID    int
	Name  string
	Price int
}

type Order struct {
	ID          int
	UserID      int
	OrderDate   string
	TotalAmount int
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
