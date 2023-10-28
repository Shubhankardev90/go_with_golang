package main

import "fmt"

func main() {

	fmt.Println("Pointers")
	// var ptr *int
	// fmt.Println("Valur of pointer is", ptr)

	myNum := 3
	var ptr = &myNum
	fmt.Println("Valur of pointer is", ptr)
	fmt.Println("Valur of pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Valur of pointer is", *ptr)

}
