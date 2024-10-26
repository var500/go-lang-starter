package greetings

import (
	"errors"
	"fmt"
);

func Hello(name string) (string, error){

// 	var message string
// 	message = fmt.Sprintf("Hi, %v. Welcome!", name)

//  The above code is written in a single line using := 
//  := operator is a shortcut for declaring and initializing a variable 	
//	in one line (Go uses the value on the right to determine the 
//	variable's type)


	// If no name was given, return an error with a message.
	if name == "" {
        return "", errors.New("empty name")
    }

	message:= fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}