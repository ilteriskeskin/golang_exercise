package main

import (
	"fmt"
	"math"
)

func main()  {
	r := 3.0

	const pi float64 = 3.14

	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)
}