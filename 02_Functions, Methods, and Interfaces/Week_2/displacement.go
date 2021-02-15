package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a float64
	var vo float64
	var so float64
	var t float64

	a = getValue("acceleration")
	vo = getValue("initial velocity")
	so = getValue("initial displacement")
	t = getValue("time")

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println("The displacement is:", fn(t))

}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5 * acceleration * math.Pow(time, 2) + (velocity * time) + displacement
	}
	return fn
}

func getValue(valType string) float64 {
	var userInput float64
	valid := false
	reader := bufio.NewScanner(os.Stdin)
	for valid == false {
		fmt.Printf("enter value for %s: ", valType)
		_ = reader.Scan()
		userInput, err := strconv.ParseFloat(reader.Text(), 64)
		if err != nil {
			fmt.Println("invalid input")
		} else {
			valid = true
			return userInput
		}
	}
	return userInput
}
