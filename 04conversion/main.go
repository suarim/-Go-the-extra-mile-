package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Your code goes here
	fmt.Println("Conversion in Go")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a number = ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
	z, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Converted value is: ", z)
	}
	fmt.Printf("type of converted %T", z)
}
