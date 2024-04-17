package Test_Tasks

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDelBoo(t *testing.T) {
	req, err := http.NewRequest("DELETE", "https://restful-booker.herokuapp.com/booking/2", nil)
	if err != nil {
		fmt.Println("Error Get request:", err)
	}
	req.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQxMjM=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Error Deleting:", resp.StatusCode)
	} else {
		fmt.Println("Success Deleting:", resp.Proto, resp.StatusCode)
	}
	defer resp.Body.Close()

}
