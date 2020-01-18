package main

import "fmt"

// x:=43  short operator will not work outside func
var y = "Hello"
var z int //to declare variable with identifier int

func main() {
	var x = 42
	fmt.Println(x)
	fmt.Println("z =>", z)

	foo()
}

func foo() {
	fmt.Println(y)
}
