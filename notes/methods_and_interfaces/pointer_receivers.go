package methods_and_interfaces

import "math"

// the main takeaway here is that
// when you declare a method on a type/struct
// you can either pass it as a refference or as a value
// when you use the pointer syntax the method will directly modify the
// fields on that type
// On the other hand passing as value will just create a copy of the
// receiver type
type TheStruct struct {
	X, Y float64
}

func (v TheStruct) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *TheStruct) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
