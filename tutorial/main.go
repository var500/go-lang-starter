package main

import "fmt"

func main() {
	fmt.Println(("Hello World"))

	//strings
	var nameOne string = "Varun"

	//auto data type detected
	var nameTwo = "Varun"

	// variable declaration default initialization
	var nameThree string

	nameOne = "Peach"
	nameTwo = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	//  short hand variable declaration and initialization
	nameFour := "Youshi"

	fmt.Println(nameFour)

	//Integerts

	var ageOne = 20
	var ageTwo int = 25
	var ageThree int // default 0
	ageFour := 30

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	//bits and memeory - 8, 16,32, 64
	// 8 bit >>  -128 to 127
	// var numOne int8 = -128
	// var numTwo uint8 = 255

	// floats
	var scoreOne float32 = 3.14159
	var scoreTwo float64 = 7127827.1221
	scoreThree := 887.988

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
