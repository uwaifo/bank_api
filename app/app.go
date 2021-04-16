package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//Optional to use mux rather than bthe default http multiplexer

	router := mux.NewRouter()

	// Define our routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/oldcustomers", getXMLCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// starting our server
	log.Fatal(http.ListenAndServe("localhost:8400", router))

}
