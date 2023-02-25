package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(), "first abs")
	v.Scale(10)
	fmt.Println(v.Abs(),"second abs")
	q := &Vertex{3, 4} // this is a pointer to vertex
	Scale(q, 10) // function scale 
	fmt.Println(q.Abs(),"third abs")

	// when using a method you can use &q or just q and it works 
}
