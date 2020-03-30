package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Printf("Exiting...")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel: ", v)

		case v := <-o:
			fmt.Println("from the odd channel: ", v)

		case v := <-q:
			fmt.Println("from the quit channel: ", v)
			return
		}
		return
	}
}

//send
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
