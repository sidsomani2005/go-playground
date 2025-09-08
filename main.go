package main

import (
	"fmt"
	"strconv"
)

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

	fmt.Println(sampleFunction("Sid is of age ", 20))
	anonymousFunction()

	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")

	fmt.Println("Doing some work...")

	fmt.Println(personalInfo(Person{Firstname: "Sid", Lastname: "Somani", Age: 20}))
}

func sampleFunction(p1 string, p2 int) string {
	return p1 + strconv.Itoa(p2)
}

func anonymousFunction() {
	func() {
		fmt.Println("I am an anonymous function")
	}()
}

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func personalInfo(p Person) string {
	return fmt.Sprintf("%s %s is %d years old", p.Firstname, p.Lastname, p.Age)
}
