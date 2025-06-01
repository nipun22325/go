package main

import (
    "fmt"
    "log"
    "example.com/randomGreetings"
)

func main() {
    log.SetPrefix("randomGreetings: ")
    log.SetFlags(0)
   
    message, err := randomGreetings.Hello("Nipun")
    if err != nil {
        log.Fatal(err)          
    }
    
    fmt.Println(message)
}