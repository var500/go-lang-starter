package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func circleArea(r float64) float64 {
	area := math.Pi * math.Pow(r, 2)
	return area
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	// n := "Mario"
	// sayGreeting(n)
	// sayBye(n)

	cycleNames([]string{"Cloud", "Barret", "tifa"}, sayGreeting)

	a1 := circleArea(10.5)
	fmt.Printf("Area of circle with radius %v is: %0.3f", 10.5, a1)
}
