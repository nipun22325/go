package main

import (
    "fmt"
    "log"
    "example.com/greetingsErr"
)

func main(){
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetingsErr: ")
    log.SetFlags(0)
    
    // Request a greeting message.
    message, err := greetingsErr.Hello("")
    if err != nil {
        log.Fatal(err) // if an error was returned, print it to the console and exit the program
    }
    
    //if no error was returned , print the returned greeting
    fmt.Println(message)
}
