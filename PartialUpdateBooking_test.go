package Test_Tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestPartUpdBoo(t *testing.T) {
	book := map[string]interface{}{
		"firstname": "Pendalf",
		"lastname":  "Seryi",
	}

	reqBody, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Can`t Marshal to JSON:", err)
	}

	req, err := http.NewRequest("PATCH", "https://restful-booker.herokuapp.com/booking/1", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error Get request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQxMjM=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error to HTTP request:", resp.StatusCode)
	} else {
		fmt.Println("Success HTTP request:", resp.Proto, resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Host reply:", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error JSON unmarshal:", err)
	}
	jsonForm, err := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonForm), "\n")

}
