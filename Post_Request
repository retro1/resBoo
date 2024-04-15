package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	//Инициализация переменной с конвертацией структуры в формат JSON
	reqBody, err := json.Marshal(map[string]string{
		"username": "admin",
		"password": "password123",
	})
	if err != nil {
		fmt.Println("Error to JSON encode:", err)
	}

	//Создание POST запроса с указанием URL, content-type и 'содержимое буфера?' или 'байтового среза?' reqBody
	req, err := http.Post("https://restful-booker.herokuapp.com/auth", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error POST request:", err)
		return
	} else {
		fmt.Println("POST request successfully")
	}

	//Гарантированное завершение функции request
	defer req.Body.Close()

	//Чтение ответа сервера
	rep, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error Host reply:", err)
	}
	fmt.Println(string(rep))

}
