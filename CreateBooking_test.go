package Test_Tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type Booking struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Totalprice   int    `json:"totalprice"`
	Depositpaid  bool   `json:"depositpaid"`
	Bookingdates struct {
		Checkin  string `json:"checkin"`
		Checkout string `json:"checkout"`
	} `json:"bookingdates"`
	Additionalneeds string `json:"additionalneeds"`
}

func TestCreateBoo(t *testing.T) {
	book := Booking{
		Firstname:   "Sasha",
		Lastname:    "Belyi",
		Totalprice:  777,
		Depositpaid: true,
		Bookingdates: struct {
			Checkin  string `json:"checkin"`
			Checkout string `json:"checkout"`
		}{
			Checkin:  "2018-01-01",
			Checkout: "2019-01-01",
		},
		Additionalneeds: "Breakfast",
	}

	reqBody, err := json.Marshal(book)
	if err != nil {
		t.Error("Can`t Marshal to JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://restful-booker.herokuapp.com/booking", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error("Error Get request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Error("Error to HTTP request:", resp.StatusCode)
		return
	} else {
		t.Log("Success HTTP request:", resp.Proto, resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error Host reply:", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Error("Error JSON unmarshal:", err)
		return
	}
	jsonForm, err := json.MarshalIndent(data, "", "\t")
	fmt.Printf("Received JSON:\n%s%s", string(jsonForm), "\n")

}
