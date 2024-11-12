package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Book struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func fetchBooksData(endpoint string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания GET-запроса: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения GET-запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения: %w", err)
	}

	return body, nil
}

func parseBooksData(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("ошибка десериализации: %w", err)
	}
	return books, nil
}

func createBook(endpoint string, data []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания POST-запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка POST-запроса:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return nil, err
	}
	return body, nil
}

const baseURL = "http://localhost:8080"

func main() {
	createBookURL := baseURL + "/createBook"

	newBook := Book{
		ID:     2017,
		Name:   "JuniorHunter",
		Author: "Saveliy",
	}

	postData, err := json.Marshal(newBook)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	responseData, err := createBook(createBookURL, postData)
	if err != nil {
		fmt.Println("Ошибка POST-запроса:", err)
		return
	}

	fmt.Printf("Response data: %s\n", responseData)

	getBooksURL := baseURL + "/getBooks"

	data, err := fetchBooksData(getBooksURL)
	if err != nil {
		fmt.Println("Ошибка GET-запроса:", err)
		return
	}

	books, err := parseBooksData(data)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	fmt.Printf("Books: %+v\n", books)
}
