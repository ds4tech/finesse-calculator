package main

import (
	"fmt"
	"log"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ds4tech/pipeline-calculator-ws/calculator_ws"
)
var dat map[string]interface{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Calculator!")
}
func logHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseFloat(num, 64)
  result := calc.Log(i)
	if err == nil {
    jsonMap := map[string]float64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
    fmt.Printf("\nLn of %v is: %v", num, result)
		fmt.Fprintf(w, string(jsonResult)) 
	}
}
func sqrtHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseFloat(num, 64)
  result := calc.Sqrt(i)
	if err == nil {
    jsonMap := map[string]float64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
    fmt.Printf("\nSqrt of %v is: %v", num, result)
		fmt.Fprintf(w, string(jsonResult)) 
	}
}


func main() {
	fmt.Printf("Sum of %d and %d is %d\n", 3,4, calc.Sum(3,4))
	fmt.Printf("Factorial of 4 is %v\n", calc.Factorial(4))
  fmt.Printf("4 is prime: %v\n", calc.IsPrime(4))
  fmt.Printf("3 is prime: %v\n", calc.IsPrime(3))
	fmt.Println("End of calculations")

	fmt.Println("Starting server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/sqrt", sqrtHandler).Methods("POST")
	router.HandleFunc("/log", logHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
