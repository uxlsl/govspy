package main

import "fmt"

func main() {
	numbers := [5]int{0, 0, 0, 0, 0}
	numbers[2] = 100
	some_numbers := numbers[1:3]
	fmt.Println(some_numbers)
	fmt.Println(len(numbers))

	scores := make([]float32, 0)
	scores = append(scores, 1.1)
	scores[0] = 2.2
	fmt.Println(scores)
}
