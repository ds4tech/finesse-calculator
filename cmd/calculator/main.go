package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/pipeline-calculator-ws/pkg"
)


func main() {
	fmt.Println("Starting server...\n")

	router := calc.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
