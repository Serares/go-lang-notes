package basics

import (
	"fmt"
	"math"
	"strings"
)

// struct
type Vertex struct {
	X int
	Y int
}

type GeoLoc struct {
	Lat, Long float64
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

func CreateMap() {
	var m map[string]GeoLoc
	m = make(map[string]GeoLoc)
	m["Bell Labs"] = GeoLoc{
		40.68433, -74.39967,
	}
	// creating a map literal

	var w map[string]GeoLoc = map[string]GeoLoc{
		// you can omit the 'GeoLoc' if it's just a type name
		"Google": GeoLoc{65.23, 11.32},
	}

	fmt.Println(w["Google"])
	fmt.Println(m["Bell Labs"])
}

func WordCount(s string) map[string]int {
	var words_list []string = strings.Split(s, " ")
	var count_words map[string]int = make(map[string]int)
	for _, word := range words_list {
		if _, ok := count_words[word]; ok != false {
			count_words[word] += 1
		} else {
			count_words[word] = 1
		}
	}

	return count_words
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunctionValues() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func FiboWithClosure() func() int {
	var fibo_list []int = []int{0, 1}
	return func() int {
		fibo_list = append(fibo_list, fibo_list[len(fibo_list)-1]+fibo_list[len(fibo_list)-2])
		return fibo_list[len(fibo_list)-1]
	}
}
