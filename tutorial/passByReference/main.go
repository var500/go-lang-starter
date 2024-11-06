// By default go creates a copy of a variable and updates the temp variable value not the real value i.e. Pass By Value

package main

import "fmt"

func updateName(n string) string {
	n = "Mario"
	return n
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	name := "Varun"

	name = updateName(name)

	fmt.Println(name)

	// Group A: int, string, bool. floats, arrays, structs
	// These data types are handled as pass my value

	//Group B:
	// type -> slices, map, functions
	// These datatypes are handled as pass by reference where the copy of variable points to the same location where data is stored

	menu := map[string]float64{
		"pie":     8.99,
		"chicken": 12.99,
	}

	fmt.Println(menu)

	updateMenu(menu)

	fmt.Println(menu)
}
