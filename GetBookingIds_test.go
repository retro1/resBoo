package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetBook(t *testing.T) {
	resp, err := http.Get("https://restful-booker.herokuapp.com/booking")
	if err != nil {
		fmt.Println("Error Get response:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read response:", err)
	}
	fmt.Println(string(body))
}
