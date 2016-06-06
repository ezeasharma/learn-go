package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
     "log"
)

//HandleQuoteRequest will be responsible for handling all the web requests.
func HandleQuoteRequest(responseWrite http.ResponseWriter, request *http.Request)  {
    log.Println("Got request!")
    resp, err := http.Get("http://inspireme.ranjithzachariah.com")
    //resp, err := http.Get("https://google.com")
    if err != nil {
        fmt.Fprintf(responseWrite, "Error connecting to InspireMe server %s", err)
    }  else {
        htmlData, _ := ioutil.ReadAll(resp.Body)
        fmt.Fprintf(responseWrite, string(htmlData))   
    }
}