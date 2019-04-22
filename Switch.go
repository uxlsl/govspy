package main

import "os"
import "bufio"
import "fmt"
import "strconv"
import "strings"

func input_() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	v, err := strconv.Atoi(strings.Trim(text, "\n"))
	if err == nil {
		return v
	} else {
		return 0
	}
}

func main() {
	number := input_()
	switch number {
	case 8:
		fmt.Println("Oxygen")
	case 1:
		fmt.Println("Hydrogen")
	case 2:
		fmt.Println("Helium")
	case 11:
		fmt.Println("Sodium")
	default:
		fmt.Printf("I have no idea what %d is", number)
	}
}
