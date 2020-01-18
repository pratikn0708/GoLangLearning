package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Ketty": 28}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["test"]) // output is zero as no key with name test is there

	v, ok := m["test"] // to check value exist or not in map
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("This is if statement", v)
	}

	// add element to map
	m["Thomas"] = 55
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete an element in map

	delete(m, "James")
	fmt.Println(m)
}
