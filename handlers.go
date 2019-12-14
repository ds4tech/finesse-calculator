package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"

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
    fmt.Printf("Sum of %v and %v is: %v\n", n1, n2, result)
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
    fmt.Printf("Sqrt of %v is: %v\n", num, result)
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
    fmt.Printf("Ln of %v is: %v\n", num, result)
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
    fmt.Printf("Factorial of %v is: %v\n", num, result)
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
	  fmt.Printf("%v is prime: %v\n", num, result)
	}
}
