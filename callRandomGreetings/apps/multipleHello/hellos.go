package main

import (
    "fmt"
    "log"
    "example.com/randomGreetings"
)

func main() {
    log.SetPrefix("randomGreetings: ")
    log.SetFlags(0)
    
    // slice of names
    names := []string{"Nipun", "Akriti", "Suhani", "Rahul"}
    messages, err := randomGreetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
    for _,greeting := range messages {
        fmt.Println(greeting)
    }
}