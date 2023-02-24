package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
var c, python, java bool
var i int
var k, j int = 1, 2
var b, sharp, dart = true, false, "no!"



func main() {
	s, d, r := true, 40, "first"
	fmt.Println(s, d, r)
	fmt.Println(i, c, python, java)   //false, false, false (default bool value)
	fmt.Println(k, j, b, sharp, dart) // true, false, "no!"
	fmt.Println(math.Pi)
	fmt.Println("My favorite number is ", rand.Intn(20))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Println(split(17))

	//Flows
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	//pointers
	i := 42
	fmt.Println("first value for i", i)
	p := &i //p is the location where "i" is saved ==0x14000126068
	fmt.Println("pointer ", p)
	*p = 21 //sets the value of i to 21
	fmt.Println("second value for i", i)
	fmt.Println("pointer ", *p) // with *p you get the value stored in that spot
	v := Vertex{1, 2}
	q := &v
	q.X = 1e9
	fmt.Println(v)
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		w  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, w, v2, v3)

	n := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(n)
}
// /A return statement without arguments returns the named return values.
// This is known as a "naked" return.
// /Naked return statements should be used only in short functions, as with the example shown here.
// /They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type Vertex struct {
	X int
	Y int
}
