package main

import "fmt"

func main() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var y []string
	if len(y) != 0 {
		fmt.Println("this won't be printed")
	}
}
