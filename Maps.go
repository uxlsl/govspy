package main

import "fmt"

func main() {
	var elements = make(map[string]int)
	elements["H"] = 1
	fmt.Println(elements["H"])
	elements["O"] = 8
	delete(elements, "O")

	if _, ok := elements["O"]; ok {
		fmt.Println(elements["O"])
	}

	if _, ok := elements["H"]; ok {
		fmt.Println(elements["H"])
	}
}
