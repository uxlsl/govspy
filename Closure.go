package main

import "fmt"

func main() {
	number := 0
	increment := func(amout int) {
		number += amout
	}
	increment(1)
	increment(2)
	fmt.Println(number)
}
