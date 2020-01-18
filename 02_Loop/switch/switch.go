package main

import "fmt"

func main() {
	a := 2
	switch a {
	case 2:
		fmt.Println("Its 2")
	case 4:
		fmt.Println("Its 4")
	default:
		fmt.Println("Its ", a)
	}

}
