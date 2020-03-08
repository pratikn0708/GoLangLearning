package main

import "fmt"

func main() {
	x := foo()
	i := x()
	fmt.Println(i)

	f := func() {
		fmt.Println("We are in Anonymous Function")
	}

	f()

	func(x int) {
		fmt.Println("The Lucky no is: ", x)
	}(22)

}

func foo() func() int {
	fmt.Println("Hello We are in foo Function")
	return func() int {
		return 2000
	}
}
