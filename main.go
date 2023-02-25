package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
//type switch cases 
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	//If i holds a T, then t will be the underlying value and ok will be true.
	s, ok := i.(string)
	fmt.Println(s, ok)

	w, ok := i.(float64)
	fmt.Println(w, ok)

	w = i.(float64) // panic
	fmt.Println(w)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
