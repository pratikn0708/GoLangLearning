package main

import "fmt"

// 2D slice
func main() {
	a := []string{"James", "Bond"}
	b := []string{"Gun", "Spy"}

	fmt.Println(a)
	fmt.Println(b)

	ab := [][]string{a, b}
	fmt.Println(ab)
}
