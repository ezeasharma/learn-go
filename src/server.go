package main

import (
    "net/http"
    "log"
)

func main(){
    log.Printf("listening on 8080")
    http.HandleFunc("/quote", HandleQuoteRequest)
    err :=http.ListenAndServe(":8080", nil)
    log.Fatal(err)
}


