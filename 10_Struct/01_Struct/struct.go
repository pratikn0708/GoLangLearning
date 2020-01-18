package main

import "fmt"

// struct is an aggregate datatype
type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond"}

	p2 := person{
		firstName: "Harry",
		lastName:  "Gill"}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.lastName)
	fmt.Println(p2.lastName)
}
