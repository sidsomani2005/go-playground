package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	if x := 10; x > 5 {
		fmt.Println("x is gt 5")
	} else {
		fmt.Println("x is not gt 5")
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
