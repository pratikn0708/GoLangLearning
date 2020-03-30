package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Printf("Exiting...")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
