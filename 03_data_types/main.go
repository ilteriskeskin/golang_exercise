package main

import "fmt"

func main()  {
	name := "İlteriş"
	age := 23
	isMarried := false
	height := 1.79
	var variable1 uint8 = 255 // uint8 0-255
	var variable2 uint16 = 65535 // 0-65535

	// more information https://golang.org/ref/spec#Types

	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", age)
	fmt.Printf("%T \n", isMarried)
	fmt.Printf("%T \n", height)
	fmt.Printf("%T \n",variable1)
	fmt.Printf("%T \n",variable2)
}