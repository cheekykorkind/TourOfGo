package main

import (
	"fmt"
	"math"
)

// var x, y, z int = 1, 2, 3
func pow (x, n, lim float64) float64 {
	// if var v float64 = math.Pow(x, n); v < lim {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}