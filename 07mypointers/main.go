// Correcting the assignment to ptr3
package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var ptr *int
	fmt.Println("The value of ptr is:", ptr)

	number := 23
	var ptr2 = &number
	fmt.Println("The value of ptr2 is:", ptr2)
	fmt.Println("The actual value of ptr2 is:", *ptr2)

	*ptr2 = *ptr2 * 5
	fmt.Println("The changed value address of ptr2 is:", ptr2)
	fmt.Println("The actual changed value of ptr2 is:", *ptr2)

	// Correcting the assignment to ptr3
	var ptr3 = ptr2
	fmt.Println("The value address of ptr3 is:", ptr3)
	fmt.Println("The actual value of ptr3 is:", *ptr3)

	*ptr3 = 78
	fmt.Println("The changed value address of ptr3 is:", ptr3)
	fmt.Println("The actual changed value of ptr2 is:", *ptr2)
}
