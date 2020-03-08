package main

import "fmt"

// struct is an aggregate datatype
type person struct {
	firstName string
	lastName  string
}
type secretAgent struct {
	person // embedd one type to another struct
	ltk    bool
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called human")
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
	}

	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond"},
		ltk: true}

	fmt.Println(p1)
	fmt.Println(sa1)
	fmt.Println(sa1.firstName, sa1.lastName, sa1.ltk)

}
