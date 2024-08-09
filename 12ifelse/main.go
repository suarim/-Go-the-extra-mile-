package main

import "fmt"

func main() {
	// Your code goes here
	fmt.Println("Hello, world!")
	x := 10
	y := 10
	if x < y {
		fmt.Println("x is less than y")
	} else if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is greater than y")
	}
	if x, y := 1, 2; x < y {
		fmt.Println("x is less than y")
	}

}
