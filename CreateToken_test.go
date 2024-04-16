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
	reqBody, err := json.Marshal(map[string]string{
		"username": "admin",
		"password": "password123",
	})
	if err != nil {
		fmt.Println("Error to JSON encode:", err)
	}

	req, err := http.Post("https://restful-booker.herokuapp.com/auth", "application/json", bytes.NewBuffer(reqBody))
	if req.StatusCode != http.StatusOK {
		fmt.Println("Error to HTTP request:", req.StatusCode)
	} else {
		fmt.Println("Success HTTP request:", req.Proto, req.Status)
	}
	defer req.Body.Close()

	rep, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error Host reply:", err)
	}
	fmt.Println(string(rep))

}
