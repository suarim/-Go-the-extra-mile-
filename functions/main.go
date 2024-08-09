package main

import (
	"bufio"
	"fmt"
	"os"
)

func greet(x int, y int, z string) {
	fmt.Println(x + y)
	fmt.Println(z)
}

func proadder(val ...int) int {
	s := 0
	for _, i := range val {
		s += i
	}
	return s
}

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	x, _ := reader.ReadString('\n')
	greet(7, 8, x)
	y := proadder(1, 2, 3, 4, 5)
	fmt.Println(y)
}
