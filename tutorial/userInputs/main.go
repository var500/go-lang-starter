package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Helper Function

func getInput(prompt string, reader *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Customer Name: ", reader)

	b := newBill(name)
	fmt.Println("Created Bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose Option: (a - add item, s- save bill, t - add tip): ", reader)

	fmt.Println(option)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)
}
