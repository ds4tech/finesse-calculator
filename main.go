package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	fmt.Println("Starting server...\n")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/sum", sumHandler).Methods("POST")
	router.HandleFunc("/sqrt", sqrtHandler).Methods("POST")
	router.HandleFunc("/log", logHandler).Methods("POST")
	router.HandleFunc("/factorial", factorialHandler).Methods("POST")
	router.HandleFunc("/isPrime", isPrimeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
