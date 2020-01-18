package main

import "fmt"

func main() {
	s := "Pratik Garg"
	b := []byte(s)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	for index := 0; index < len(s); index++ {
		fmt.Printf("%#U ", s[index]) //UTFA code
	}
}
