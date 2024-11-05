package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"chicken curry": 20.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["chicken curry"])

	for k, v := range menu {
		fmt.Println(k, " - ", v)
	}

	//ints as key type
	phoneBook := map[int]string{
		123456: "Mario",
		234567: "Luigi",
		345678: "Peach",
	}

	fmt.Println(phoneBook[123456])
	for k, v := range phoneBook {
		fmt.Println(k, " - ", v)
	}

	phoneBook[123456] = "Bowser"

	for k, v := range phoneBook {
		fmt.Println(k, " - ", v)
	}
}
