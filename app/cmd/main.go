package main

import (
    "fmt"
    "add"
)

func main() {
    fmt.Println("hello go rajesh")
    addResult := add.Add(5, 3)
    fmt.Printf("The result of addition is: %d\n", addResult)
}
