package main

import (
    "fmt"
    "log"
)

func main() {

    user := make(map[string]interface{}, 0)

    user["name"] = "John Doe"
    user["age"] = 21
    user["weight"] = 70.3

    age, ok := user["age"].(int)

    if !ok {
        log.Fatal("assert failed")
    }

    user["age"] = age + 1

    fmt.Printf("%+v", user)
}