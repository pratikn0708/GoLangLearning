package main

import "fmt"

var y = 42
var z = "Go Language is Fun"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// z = 43 cant assign int to string as go Lang is static language and is purely type based lang

	// go is static programming lang not dynamic like JS which is dynamic lang
}
