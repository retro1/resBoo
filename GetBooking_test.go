package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetBooking(t *testing.T) {
	req, err := http.NewRequest("GET", "https://restful-booker.herokuapp.com/booking/3", nil)
	if err != nil {
		fmt.Println("Error Get request:", err)
	}

	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error Get request:", resp.StatusCode)
	} else {
		fmt.Println("Success:", resp.Proto, resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error JSON unmarshal:", err)
	}
	jsonF, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonF), "\n")
}
