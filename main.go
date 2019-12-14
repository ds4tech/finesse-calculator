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

func sumHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer numbers")
	}
  json.Unmarshal(reqBody, &dat)
  n1 := dat["num1"].(string)
  n2 := dat["num2"].(string)

  i1, err := strconv.ParseInt(n1, 10, 64)
	i2, err := strconv.ParseInt(n2, 10, 64)
  result := calc.Sum(i1, i2)
	if err == nil {
    jsonMap := map[string]int64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
    fmt.Printf("\nSum of %v and %v is: %v", n1, n2, result)
	}
}

func sqrtHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseFloat(num, 64)
  result := calc.Sqrt(i)
	if err == nil {
    jsonMap := map[string]float64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
    fmt.Printf("\nSqrt of %v is: %v", num, result)
	}
}
func logHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseFloat(num, 64)
  result := calc.Log(i)
	if err == nil {
    jsonMap := map[string]float64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
    fmt.Printf("\nLn of %v is: %v", num, result)
	}
}
func factorialHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseUint(num, 10, 64)
  result := calc.Factorial(i)
	if err == nil {
    jsonMap := map[string]uint64{"result": result}
    jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
    fmt.Printf("\nFactorial of %v is: %v", num, result)
	}
}
func isPrimeHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
  json.Unmarshal(reqBody, &dat)
  num := dat["number"].(string)

  i, err := strconv.ParseInt(num, 10, 64)
  result := calc.IsPrime(int(i))
	if err == nil {
    jsonMap := map[string]bool{"isPrime": result}
    jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
	  fmt.Printf("\n%v is prime: %v\n", num, result)
	}
}

func main() {
	fmt.Printf("Sum of %d and %d is %d\n", 3,4, calc.Sum(3,4))
  fmt.Printf("4 is prime: %v\n", calc.IsPrime(4))
  fmt.Printf("3 is prime: %v\n", calc.IsPrime(3))
	fmt.Println("End of calculations")

	fmt.Println("Starting server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/sum", sumHandler).Methods("POST")
	router.HandleFunc("/sqrt", sqrtHandler).Methods("POST")
	router.HandleFunc("/log", logHandler).Methods("POST")
	router.HandleFunc("/factorial", factorialHandler).Methods("POST")
	router.HandleFunc("/isPrime", isPrimeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
