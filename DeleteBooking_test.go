package Test_Tasks

import (
	"net/http"
	"testing"
)

func TestDelBoo(t *testing.T) {
	req, err := http.NewRequest("DELETE", "https://restful-booker.herokuapp.com/booking/1", nil)
	if err != nil {
		t.Error("Error Get request:", err)
		return
	}
	req.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQxMjM=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusCreated {
		t.Error("Error Deleting:", resp.StatusCode)
		return
	} else {
		t.Log("Success Deleting:", resp.Proto, resp.StatusCode)
	}
	defer resp.Body.Close()

}
