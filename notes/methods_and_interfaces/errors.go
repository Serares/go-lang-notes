package methods_and_interfaces

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// convert e to a float or the Error will be called in an infinite loop
	return fmt.Sprintf("You can't sqrt this number %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x / 2.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func HandleErrors() {
	fmt.Println(Sqrt(-2))
}
