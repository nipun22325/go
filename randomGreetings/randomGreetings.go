package randomGreetings

import (
    "fmt"
    "errors"
    "math/rand" // to return a random greeting
)

func Hello(name string) (string, error) {
    if name == "" {
        return name, errors.New("empty name")
    }
    // create a message using a random format
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
// randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
    // A slice of message formats.
    // A slice is like an array, except that its size changes dynamically as you add and remove items.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    // Return a randomly selected message format by specifying a random index for the slice of formats.
    // Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <= 0.
    return formats[rand.Intn(len(formats))]
}