package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)  // type
	fmt.Printf("%b\n", x)  // binary
	fmt.Printf("%x\n", x)  // hexadecimal
	fmt.Printf("%#x\n", x) // hexadecimal with zero
	x = 911
	fmt.Printf("%#x\n", x)
	fmt.Printf("%#x\t%b\t%x\n", x, x, x)

	s := fmt.Sprintf("%#x\t%b\t%x\n", x, x, x)
	fmt.Println(s) // this is to print in string format

	fmt.Printf("%v\n", x) // value in default format

}
