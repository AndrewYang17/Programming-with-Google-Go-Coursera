package main

import (
	"fmt"
	"math"
)

func main() {
	dist1 := makeDistOrigin(0, 0)
	dist2 := makeDistOrigin(2, 2)
	fmt.Println(dist1(2, 2))
	fmt.Println(dist2(2, 2))
}

func makeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}
