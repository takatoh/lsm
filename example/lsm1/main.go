package main

import (
	"fmt"

	"github.com/takatoh/lsm"
)

func main() {
	var xs []float64 = []float64{ 0.0, 0.2, 0.4, 0.6, 0.8, 1.0, 1.2 }
	var ys []float64 = []float64{ 1.0, 1.9, 3.2, 4.3, 4.8, 6.1, 7.2 }

	a, b := lsm.LSM1(xs, ys)

	fmt.Printf("a = %f, b = %f\n", a, b)
}
