package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// will return a greetings for the named person
func Hello(name string) (string, error) {

	// handling some errors
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting that embeds the name in a message
	//txt := fmt.Sprintf(randomFormat(), name)
	txt := fmt.Sprint(randomFormat())
	return txt, nil
}

/*
	hellos returns a map that that associates the named people
	with a greetings messaage
*/

func Hellos(names []string) (map[string]string, error) {
	// a map to associate the names with messages
	texts := make(map[string]string)
	//loop thorough the received slice of names, calling
	// the hello function to get a messaage for each name
	for _, name := range names {
		txt, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved txt with the name
		texts[name] = txt
	}
	return texts, nil
}

/*
	randomFormat returns one of a set of greetings message. the returned
	messafe is selected randomly
*/

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
