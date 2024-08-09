package main

import "fmt"

const Constant int = 456789

func main() {
	// Your code here
	//explicit declaration of variables
	var username string = "John Doe"
	var isbool bool = true
	var isint uint64 = 100
	var x = 5
	fmt.Println(x)
	var osfloat float64 = 670.0
	fmt.Println("Variables in Go")
	fmt.Println(username)
	fmt.Printf("variable is of type: %T\n", username)
	fmt.Printf("variable is of type: %T\n", isint)
	fmt.Printf("variable is of type: %T\n", isbool)
	fmt.Printf("variable is of type: %T\n", osfloat)

	//implicit declaration of variables
	var username2 = "John Doe"
	fmt.Println(username2)

	//novar declaration of variables
	age := 20
	fmt.Println(age)

	fmt.Println(Constant)
}
