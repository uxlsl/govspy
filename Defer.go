package main

import "os"

func main(){
	f, _ := os.Open("defer.py")
	defer f.Close()
}
