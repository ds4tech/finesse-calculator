package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	fmt.Println("Starting server...\n")

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
