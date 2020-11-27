package main

import "fmt"

func main()  {
	var (
		name string = "İlteriş"
		age  int    = 23
	)

	// var name, age = "İlteriş", 23
	// name, age := "İlteriş", 23

	fmt.Println(name, age)

	var zeroValue1 string
	var zeroValue2 int
	var zeroValue3 bool

	fmt.Println(zeroValue1, zeroValue2, zeroValue3)
}