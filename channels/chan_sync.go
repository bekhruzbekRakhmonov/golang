package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("something went wrong")

    done <- false
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    bool := <-done

    fmt.Println(bool)
}