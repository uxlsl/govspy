package main

import "fmt"
import "time"

func main() {
	t0 := time.Now()
	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Printf("Took %s seconds", t1.Sub(t0))
}
