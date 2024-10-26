package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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
        return name, errors.New("empty name")
    }

	message:= fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

	return formats[rand.Intn((len(formats)))]
}

func Hellos(names []string) (map[string]string, error){
	messages:= make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if(err!=nil){
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
