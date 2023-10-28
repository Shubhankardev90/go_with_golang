package main

import "fmt"

func main() {
	fmt.Println("Structs")

	//no inheritence , no Super or Parent

	shubhankar := User{"Shubhankar", "shubhankar@intellohire.com", true, 22}
	shubhankar.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is User active: ", u.Status)
}
