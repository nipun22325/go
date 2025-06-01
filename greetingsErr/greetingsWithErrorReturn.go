package greetingsErr

import (
    "errors"
    "fmt"
)

// Any Go function can return multiple values.
func Hello(name string) (string, error) {
    // If no name was given, return an error with message
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a value that embeds the name in a greeting message
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
