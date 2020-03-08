package main

import "fmt"

func main() {
	p1 := struct { // this is anonymus struct as we dont give any type name, we simply define variable and used
		firstName string
		lastName  string
	}{
		firstName: "James",
		lastName:  "Bond"}

	fmt.Println(p1)

}
