package main

import "fmt"

func main() {
	x := 10
	y := 10.5

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)


	// mismatched types int and float64
	// fmt.Println(x + y)
	fmt.Println(x + int(y))
	fmt.Println(float64(x) + y)

	var a int8 = 127
	var b int16

	b = int16(a)
	fmt.Println(b)
}