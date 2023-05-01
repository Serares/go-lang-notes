package main

import (
	"fmt"
	"runtime"

	"example.com/notes/basics"
	methods "example.com/notes/methods_and_interfaces"
)

func switchCase(n int) {
	switch n {
	case 1:
		fmt.Println("You chose death")
	case 2:
		fmt.Println("You can live")
	default:
		fmt.Println("You have to roll again")
	}
}

func main() {
	var v basics.Vertex
	fmt.Println("Hello")
	fmt.Println(runtime.GOOS)
	// switchCase(1)

	// calling the deffer example
	// basics.F()

	// Struct
	v = basics.Vertex{1, 2}
	vertexPointer := &v
	vertexPointer.X = 23 // both this and the bellow syntax are valid
	// because the compiler knows that you are trying to use a pointer to X
	// when refferencing a struct field
	// (*vertexPointer).X = 23
	fmt.Println(v.X)

	// arrays
	// basics.Arrays()

	// slices
	basics.Pic(2, 3)

	// maps
	basics.CreateMap()

	fmt.Println(basics.WordCount("this is a big string"))

	// functions
	basics.FunctionValues()
	fmt.Println(basics.FiboWithClosure())

	// Methods
	fmt.Println(methods.UseMethod())

	// Methods on any types
	f := methods.MyFloat(3)
	fmt.Println("The abs method on type MyFloat", f.Abs())

}
