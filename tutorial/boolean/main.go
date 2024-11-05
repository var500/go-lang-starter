package main

import "fmt"

func main() {
	age := 24

	fmt.Println(age <= 25)
	fmt.Println(age == 25)
	fmt.Println(age != 25)

	if age < 30 {
		fmt.Printf("Age: %v", age)
	}

	name := []string{"mario", "luigi", "bower", "peach", "youshi"}

	for index, value := range name {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("THe value at pos %v is %v \n", index, value)
	}
}
