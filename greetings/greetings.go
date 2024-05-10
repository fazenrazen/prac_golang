package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given return an error with a message.
	// error checking
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	//message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())

	return message, nil
}

// Hellos returns a map that associates each of the the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to assocaiate names with messages
	// how to make a map --> make(map[key-type]value-type)
	messages := make(map[string]string)
	// loop through the recieved slice of names, calling
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	// A slice of a message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
