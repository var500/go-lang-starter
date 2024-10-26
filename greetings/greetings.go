package greetings

import "fmt";

func Hello(name string) string{

// 	var message string
// 	message = fmt.Sprintf("Hi, %v. Welcome!", name)

//  The above code is written in a single line using := 
//  := operator is a shortcut for declaring and initializing a variable 	
//	in one line (Go uses the value on the right to determine the 
//	variable's type)

	message:= fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}