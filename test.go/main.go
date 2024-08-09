package main

import (
	// "bufio"
	"fmt"
	"sort"
	// "os"
	// "strconv"
	// "strings"
)

func main() {

	fmt.Println("dfghjh")
	var b = 67
	n := 23
	fmt.Println(n)
	fmt.Println(b)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your no.")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	// fmt.Println(strconv.ParseFloat(strings.TrimSpace(text), 64))
	ty := 90
	ptr := new(int)
	*ptr = 23
	ptr2 := new(int)
	ptr2 = &ty
	fmt.Println("The value of ptr is:", ptr)
	fmt.Println("The actual value of ptr is:", *ptr)
	fmt.Println("The value of ptr2 is:", ptr2)
	fmt.Println("The actual value of ptr2 is:", *ptr2)

	var array = []string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("The value of array is:", array)
	x := append(array, "Kiwi")
	fmt.Println("The value of array is:", array)
	fmt.Println("The value of array is:", x)
	sort.Strings(x)
	fmt.Println("The value of array is:", x)
	user := User{
		name: "Suarim",
		age:  21,
	}
	fmt.Println(user)
}

type User struct {
	name string
	age  int
}
