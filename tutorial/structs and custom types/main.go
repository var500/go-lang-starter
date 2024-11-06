package main

import "fmt"

func main() {
	myBill := newBill("Marios Bill")
	// fmt.Println(myBill)

	fmt.Println(myBill.format())
}
