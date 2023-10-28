package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
