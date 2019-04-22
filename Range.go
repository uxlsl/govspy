package main

import "fmt"

func main(){
	names := []string{"Peter", "Anders", "Bengt"}
	for i,name := range names {
		fmt.Printf("%d. %s\n", i+1,name)
	}
}
