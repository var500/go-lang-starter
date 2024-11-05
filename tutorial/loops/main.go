package main

import "fmt"

func main() {

	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of I is:", i)
	// }

	name := []string{"mario", "luigi", "bower", "peach", "youshi"}

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println("Name :", name[i])
	// }

	for index, value := range name {
		fmt.Println(index, value)
	}
}
