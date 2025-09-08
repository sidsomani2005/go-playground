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
	fmt.Println(superHeroInfo())
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

type SuperHero struct {
	Person
	Superpower string
}

func personalInfo(p Person) string {
	return fmt.Sprintf("%s %s is %d years old", p.Firstname, p.Lastname, p.Age)
}

func superHeroInfo() string {
	s := SuperHero{}
	s.Firstname = "Bruce"
	s.Lastname = "Wayne"
	s.Age = 30
	s.Superpower = "Rich"
	return fmt.Sprintf("%s %s is %d years old and has %s", s.Firstname, s.Lastname, s.Age, s.Superpower)
}
