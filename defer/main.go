package main

import "fmt"

func test() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
func test2() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}

func main() {
	defer fmt.Print("fish")
	fmt.Println("Hello, World!")
	test2()
	test()
}
