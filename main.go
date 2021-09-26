package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/solve", Solve).Methods("POST")

	port := os.Getenv("PORT")

	if port!="" {
		log.Fatal(http.ListenAndServe(":"+port, r))
	}else{
		log.Fatal(http.ListenAndServe(":9000", r))
	}
	
}

func main() {
	initializeRouter()
}