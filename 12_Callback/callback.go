package main

import (
	"fmt"
)

func main() {
	math(23, 7, add)
	math(23, 7, multiply)
}

func add(x, y int) {
	fmt.Println(x + y)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func math(x int, y int, f func(a, b int)) {
	f(x, y)
}
