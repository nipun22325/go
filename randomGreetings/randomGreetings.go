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

/* need to pass a set of names to a function that can return a greeting for each of them.
 Changing the Hello function's parameter from a single name to a set of names would change the function's 
 signature. If you had already published the example.com/greetings module and users had already written code 
 calling Hello, that change would break their programs.
 In this situation, a better choice is to write a new function with a different name. The new function will 
 take multiple parameters. That preserves the old function for backward compatibility.
*/

// Hellos returns a map that associates each of the named people with a greeting message.
// usage of map : map[key_type]value_type

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)  // a map to associate names with messages
    
    /* In this for loop, range returns two values: the index of the current item in the loop and a copy of the 
    item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.*/
    for _,name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }           
        // In the map, associate the retrieved message with the name.
        messages[name] = message
    }
    return messages, nil
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