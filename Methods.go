package main

import "math"
import "fmt"

type Point struct {
	x float64
	y float64
}

func (self Point) distance(point Point) float64 {
	return math.Sqrt(self.x*point.x + self.y*point.y)
}

func main() {
	p1 := Point{1, 3}
	p2 := Point{2, 4}
	fmt.Println(p1.distance(p2))
}
