package greetings

import "fmt"

/*
In Go, a function whose name starts with a capital letter can be called by a function not 
in the same package.
In the function below, Hello is the function name, name is a parameter followed by its type, i.e. string
and then the return type is written (here string).
*/

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name) 
    // := operator is a shortcut for declaring and initializing a variable in one line
    /* 
       the above line can also be written as:
       var message string 
       message = fmt.Sprintf("Hi, %v. Welcome!", name)
    */
    // Go uses the value on the right to determine the variable's type
    return message  
}