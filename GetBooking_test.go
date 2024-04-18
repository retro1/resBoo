package Test_Tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetBooking(t *testing.T) {
	req, err := http.NewRequest("GET", "https://restful-booker.herokuapp.com/booking/1", nil)
	if err != nil {
		t.Error("Error Get request:", err)
		return
	}

	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Error("Error Get request:", resp.StatusCode)
		return
	} else {
		t.Log("Success:", resp.Proto, resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error reading response:", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Error("Error JSON unmarshal:", err)
		return
	}
	jsonF, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonF), "\n")
}
