package main


import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

func Sum(x int, y int) int {
  return x+y
}

func Sqrt(x float64) float64 {
  return math.Sqrt(x)
}

func Factorial(n uint64) (result uint64) {
  if (n>0) {
    result = n * Factorial(n-1)
    return result
  }
  return 1
}

func isPrime(n int) (result bool) {
  for i := 2; i <= n; i++ {
    if big.NewInt(int64(i)).ProbablyPrime(20) {
      result = true
    } else {
      result = false 
    }
  }
  return result
}

func Log(n float64) (result float64) {
  result = math.Log(n)
  return result
}

func main() {
	fmt.Println(strings.ToUpper("Hello GO"))
	fmt.Printf("Sum of %d and %d is %d\n", 3,4, Sum(3,4))
	fmt.Printf("Factorial of 4 is %v\n", Factorial(4))
  fmt.Printf("4 is prime: %v\n", isPrime(4))
  fmt.Printf("3 is prime: %v\n", isPrime(3))
  fmt.Printf("Log(9) = %v\n", Log(9))
	fmt.Println("End of calculations")
}
