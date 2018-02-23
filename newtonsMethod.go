package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 100.0
	i := 0
	for i < 1 {
		j := math.Abs(z*z - x)
		if j <= .001 {
			i = 1
		} else {
			fmt.Println(z)
			z -= (z*z - x)/(2*z)
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2*math.Pi))
}
