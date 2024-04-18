package Test_Tasks

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
		t.Error("Error to HTTP request:", resp.StatusCode)
		return
	} else {
		t.Log("Success:", resp.Proto, resp.StatusCode)
	}
	defer resp.Body.Close()

	resRead, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error Read response:", err)
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal([]byte(resRead), &data)
	if err != nil {
		t.Error("Error Unmarshal:", err)
		return
	}
	jsonF, err := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonF), "\n")
}
