package Test_Tasks

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHealCheck(t *testing.T) {
	req, _ := http.Get("https://restful-booker.herokuapp.com/ping")
	if req.StatusCode != http.StatusCreated {
		fmt.Println("Error:", req.StatusCode)
	} else {
		fmt.Println("Ok", req.Proto, req.StatusCode)
	}
	defer req.Body.Close()

}
