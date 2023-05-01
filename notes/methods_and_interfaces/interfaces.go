package methods_and_interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyStruct struct {
	X, Y float64
}

func (v *MyStruct) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A value of interface type can hold any value that implements those methods.
func Interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := MyStruct{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

// How interfaces work
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func DeclareInterface() {
	var i I = T{"hello"}
	i.M()
}
