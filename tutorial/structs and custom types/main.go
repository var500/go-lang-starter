package main

import "fmt"

func main() {
	myBill := newBill("Marios Bill")
	// fmt.Println(myBill)

	// fmt.Println(myBill.format())

	myBill.updateTip(10)
	myBill.addItem("Pie", 6.99)
	myBill.addItem("Onion Soup", 4.50)
	myBill.addItem("Caesear Salad", 12.50)

	fmt.Println(myBill.format())
}
