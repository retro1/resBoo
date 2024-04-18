package Test_Tasks

import (
	"net/http"
	"testing"
)

func TestHealCheck(t *testing.T) {
	req, _ := http.Get("https://restful-booker.herokuapp.com/ping")
	if req.StatusCode != http.StatusCreated {
		t.Error("Error:", req.StatusCode)
		return
	} else {
		t.Log("Ok", req.Proto, req.StatusCode)
	}
	defer req.Body.Close()

}
