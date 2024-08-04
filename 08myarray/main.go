package main

import "fmt"

func main() {
	// Your code here
	fmt.Println("Hello World")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Orange"
	fmt.Println("The value of fruits is:", fruits)
	fmt.Println("Length = ", len(fruits))

	var veg = [4]string{"Carrot", "Potato", "Tomato", "Onion"}
	fmt.Println("The value of veg is:", veg)

}
