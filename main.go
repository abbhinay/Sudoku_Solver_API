package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/solve", Solve).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	initializeRouter()
}