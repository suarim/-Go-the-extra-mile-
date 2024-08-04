package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter value = ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
	fmt.Printf("type of input is %T\n", input)
}
