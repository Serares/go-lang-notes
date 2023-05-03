package gorutines

import "fmt"

func fibo(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func FiboGoRutines() {
	c := make(chan int, 10)
	go fibo(cap(c), c)

	// v, ok := <-ch
	// ok is false if there are no more values to receive and the channel is closed.
	for i := range c {
		fmt.Println(i)
	}
}
