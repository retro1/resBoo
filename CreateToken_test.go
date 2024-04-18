package Test_Tasks

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
		t.Error("Error to JSON encode:", err)
		return
	}

	req, err := http.Post("https://restful-booker.herokuapp.com/auth", "application/json", bytes.NewBuffer(reqBody))
	if req.StatusCode != http.StatusOK {
		t.Error("Error to HTTP request:", req.StatusCode)
		return
	} else {
		t.Log("Success HTTP request:", req.Proto, req.Status)
	}
	defer req.Body.Close()

	rep, err := io.ReadAll(req.Body)
	if err != nil {
		t.Error("Error Host reply:", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(rep, &data)
	if err != nil {
		t.Error("Error JSON unmarshal:", err)
		return
	}
	jsonForm, err := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonForm), "\n")

}
