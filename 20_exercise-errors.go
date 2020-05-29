package main

import (
	"fmt"
	"math"
)

const delta = 1e-10

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number: %v", float64(err))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	v := f
	for {
		n := v-(v*v-f)/(2*v)
		if math.Abs(n-v) < delta {
			break
		}
		v = n
	}

	return math.Floor(v), nil
}

func main() {
	fmt.Println(Sqrt(81))
	fmt.Println(Sqrt(-81))
}