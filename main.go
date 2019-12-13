package main

import (
	"github.com/ds4tech/pipeline-calculator-ws/calculator_ws"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Hello GO"))
	fmt.Printf("Sum of %d and %d is %d\n", 3,4, calc.Sum(3,4))
	fmt.Printf("Factorial of 4 is %v\n", calc.Factorial(4))
  fmt.Printf("4 is prime: %v\n", calc.IsPrime(4))
  fmt.Printf("3 is prime: %v\n", calc.IsPrime(3))
  fmt.Printf("Log(9) = %v\n", calc.Log(9))
	fmt.Println("End of calculations")
}
