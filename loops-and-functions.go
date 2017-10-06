package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 10; i ++ {
		z = z - (math.Pow(z, 2) - x) / 2 / x
		fmt.Println(z)
		if tmp := math.Pow(z, 2); math.Abs(tmp - x) < 1e-6 {
			fmt.Println(i)
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
