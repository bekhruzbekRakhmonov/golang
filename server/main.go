package main

import (
    "fmt"
    "net/http"
)


// https://golangbyexample.com/packages-modules-go-second/
// https://zetcode.com/golang/package/#:~:text=The%20package%20main%20tells%20the,entry%20point%20of%20the%20program.


func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}