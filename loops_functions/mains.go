package main

import "fmt"

func Sqrt(x float64) float64 {
	// find the most
	var z float64 = 1
	var diff float64 = 0

	for i := float64(0); i < 10; i++ {
		z += i
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
