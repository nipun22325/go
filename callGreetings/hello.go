//  In Go, code executed as an application must be in a main package.
package main

import (
    "fmt"
    "example.com/greetings"
)

func main(){
    message := greetings.Hello("Nipun")
    fmt.Println(message)
}


/*
For production use, you’d publish the example.com/greetings module from its repository
(with a module path that reflected its published location), where Go tools could find it to download it.
For now, because you haven't published the module yet, you need to adapt the example.com/hello module so 
it can find the example.com/greetings code on your local file system.
we use : go mod edit -replace example.com/greetings=../greetings
*/