package main

import "fmt"

func main() {
	l := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Hello, World!")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for _, n := range l {
		fmt.Println(n)
	}
}
