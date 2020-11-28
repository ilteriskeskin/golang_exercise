package main

import "fmt"

func main()  {
	/*
		if <boolean expression> {code}
	*/

	x := 27

	if x%2 == 0 {
		fmt.Println(x, "is even number")
	} else {
		fmt.Println(x, "is not even number")
	}
}