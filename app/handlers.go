package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"` // Optional for json naming
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip" xml:"zip"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Zen", City: "Green", Zipcode: "90076"},
		{Name: "Pipi", City: "Abuja", Zipcode: "7494"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/XML")

		//Here we do the actual encoding of the intended response payload
		xml.NewEncoder(w).Encode(customers)

	} else {
		//Line bellow handles the response writer encoding by adding the type of encoding we want to the header
		w.Header().Add("Content-Type", "application/json")

		//Here we do the actual encoding of the intended response payload
		json.NewEncoder(w).Encode(customers)

	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Post request received")

}

func getXMLCustomer(w http.ResponseWriter, r *http.Request) {
	pensioneers := []Customer{
		{Name: "Zen", City: "Green", Zipcode: "90076"},
		{Name: "Pipi", City: "Abuja", Zipcode: "7494"},
	}

	//Line bellow handles the response writer encoding by adding the type of encoding we want to the header
	w.Header().Add("Content-Type", "application/XML")

	//Here we do the actual encoding of the intended response payload
	xml.NewEncoder(w).Encode(pensioneers)

}
