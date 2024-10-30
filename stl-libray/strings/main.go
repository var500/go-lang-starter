package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "ello ther"))        //true
	fmt.Println(strings.ReplaceAll(greeting, "hello", "Yo Yo")) //true
	fmt.Println(strings.ToUpper(greeting))

}
