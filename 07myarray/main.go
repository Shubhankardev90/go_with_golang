package main

import "fmt"

func main() {

	fmt.Println("Arrays")

	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Peach"
	fruitlist[3] = "Fig"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var veglist = [3]string{"Potato", "Onion", "chilli"}
	fmt.Println(veglist)
}
