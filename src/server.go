package main

import (
    "net/http"
    "log"
)

func main(){
    log.Printf("listening on 8000")
    http.HandleFunc("/", HandleRequest)
    err :=http.ListenAndServe(":8000", nil)
    log.Fatal(err)
}


