package main

import "fmt"

func main()  {
	var name string = "Ali İlteriş"
	var surname string = "Keskin"
	// var name, surname string = "Ali İlteriş", "Keskin"
	// name := "Ali İlteriş"
	var age int = 23
	var address string = "Kale Mahallesi, 144. Sokak"
	var no int = 13
	var floor int = 1

	fmt.Println("Name   : ", name)
	fmt.Println("Surname: ", surname)
	fmt.Println("Age    : ", age)
	fmt.Println("Address: ", address, "-", no, "/", floor)
}