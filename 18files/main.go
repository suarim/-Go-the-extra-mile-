package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	content := "To be appended in file"
	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", length)
	file.Close()
	read()
}

func read() {
	// file, err := os.Open("./test.txt")
	databyte, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
	// defer file.Close()
}
