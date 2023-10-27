package main

import "fmt"

func main() {
	var username string = "Shubhankar"
	fmt.Println(username)
	fmt.Printf("Type of %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of %T", isLoggedIn)

	// implicit type
	var name = "Shubhankar"
	fmt.Println(name)

	//no var style
	noOfUser := 30000
	fmt.Println(noOfUser)

}
