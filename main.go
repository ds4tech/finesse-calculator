package main


import (
	"fmt"
	"strings"
)

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
}
