package main


import (
	"fmt"
	"math"
	"strings"
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func Foo() string {
	return "foo"
}
func Bar() string {
	return "bar"
}
func Quz(v string) string {
	if v == "foo" {
		return Foo()
	}
	if v == "bar" {
		return Bar()
	}
	return "INVALID"
}

//var b = foo()
func main() {
	fmt.Println(strings.ToUpper("Hello GO"))
  fmt.Println(sqrt(2), sqrt(-4))
}
