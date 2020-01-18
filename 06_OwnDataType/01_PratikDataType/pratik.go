package main

import "fmt"

var x int

type pratik string
type audi int

var s pratik
var y audi

func main() {

	x = 42
	s = "Pratik Garg"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// x = y  cant do that becuase x is type of int and y is type of main.audi
	//but we can do this by type conversion
	y = 69
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
