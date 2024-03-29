package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/finesse-calculator/calc"
)

func main() {
	fmt.Println("Starting server...\n")

	router := calc.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
