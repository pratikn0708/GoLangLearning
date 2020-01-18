package main

import "fmt"

// a slice allow you to group together values os same type
func main() {

	x := []int{4, 5, 6, 9}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // slicing from index 1 to 3

	x = append(x, 20) //apppend 20 at last of array
	fmt.Println(x)

	x = append(x, 40, 60, 80) //we can append mutiple value
	fmt.Println(x)

	y := []int{100, 200, 300, 400}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // delete some elements
	fmt.Println(x)

	/* make is used when there is fixed length of slice*/
	a := make([]int, 10, 100)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
}
