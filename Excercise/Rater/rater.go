package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa Center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("Hello %v \nThanks for Rating our Dosa Center with %v star rating. \n\nYour rating was recorded in our system at %v\n\n",
		name, myNumRating, time.Now().Format(time.Stamp))

	if myNumRating == 5 {
		fmt.Println("Bonus for team for 5 star service.")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("We are always improving.")
	} else if myNumRating < 3 {
		fmt.Println("Need serious work from our side.")
	}

}
