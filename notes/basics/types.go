package basics

import "fmt"

// struct
type Vertex struct {
	X int
	Y int
}

// arrays

func Arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Array"
	fmt.Println(a)

	primes := [3]int{1, 3, 5}
	fmt.Println(primes)
}

func Slices() {
	// slices are references to arrays
	// bellow is a slice literal
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// slice exercise

func Pic(dx, dy int) [][]uint8 {
	d2Slice := make([][]uint8, dy)
	for i, _ := range d2Slice {
		d2Slice[i] = make([]uint8, dx)
	}
	fmt.Println(d2Slice)
	return d2Slice
}
