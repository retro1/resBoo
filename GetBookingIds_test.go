package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetBookId(t *testing.T) {
	resp, err := http.Get("https://restful-booker.herokuapp.com/booking")
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error to HTTP request:", resp.StatusCode)
	} else {
		fmt.Println("Success:", resp.Proto, resp.StatusCode)
	}
	defer resp.Body.Close()

	resRead, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read response:", err)
	}

	var data []map[string]interface{}
	err = json.Unmarshal([]byte(resRead), &data)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
	}
	jsonF, err := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonF), "\n")
}
