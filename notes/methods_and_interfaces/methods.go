package methods_and_interfaces

import "math"

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func UseMethod() float64 {
	v := Vertex{1, 2}
	return v.Abs()
}

type TheFloat float64

func (f TheFloat) Abs() float64 {
	return float64(f)
}
