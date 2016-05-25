package main

import (
    "net/http"
    "fmt"
)

//HandleRequest will be responsible for handling all the web requests.
func HandleRequest(responseWrite http.ResponseWriter, request *http.Request)  {
    
    fmt.Printf("Got request!")
    fmt.Fprintf(responseWrite, "Hi there, I love %s!", request.URL.Path[1:])
}