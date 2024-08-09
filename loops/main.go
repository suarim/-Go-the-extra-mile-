package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for i, d := range days {
	// 	fmt.Println(i, d)
	// }
	rv := 1
	for rv < 10 {
		if rv == 5 {
			goto lco
		}
		fmt.Println(rv)
		rv++
	}

lco:
	fmt.Println("Hello, World!")
}
