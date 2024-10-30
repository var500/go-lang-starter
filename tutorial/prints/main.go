package main

import "fmt"

func main() {
	name := "Varun"
	age := 23
	city := "Kullu"
	fmt.Print("Hello")            // No line breaks
	fmt.Println("Hello Eveyrone") // Adds a line break
	fmt.Println("My age is", age, "and I live in", city, "city")
	fmt.Printf("My age is %v and my name is %q \n", age, name) // Formatted String
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("You scored %0.1f points \n", 255.55)

	// %v default format specifiers
	// %q adds double quotes arround the string
	// %T type of the variable
	// Format specifiers https://pkg.go.dev/fmt#Printf

	//Sprintf (save fromatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %q \n", age, name)

	fmt.Println("The saved string is ", str)
}
