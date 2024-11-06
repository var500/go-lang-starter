// By default go creates a copy of a variable and updates the temp variable value not the real value i.e. Pass By Value

package main

import "fmt"

func updateName(n *string) {
	*n = "Mario"
}

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

func main() {
	name := "Varun"

	// name = updateName(name)

	// fmt.Println(name)

	// Group A: Non-Pointer Values
	// type -> int, string, bool. floats, arrays, structs
	// These data types are handled as pass my value

	//Group B: Pointer Wrapper Values
	// type -> slices, map, functions
	// These datatypes are handled as pass by reference where the copy of variable points to the same location where data is stored

	// menu := map[string]float64{
	// 	"pie":     8.99,
	// 	"chicken": 12.99,
	// }

	// fmt.Println(menu)

	// updateMenu(menu)

	// fmt.Println(menu)

	// Pointers can we used like this

	m := &name
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}
