package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
     "log"
     //"encoding/json"
)

//HandleQuoteRequest will be responsible for handling all the web requests.
func HandleQuoteRequest(responseWrite http.ResponseWriter, request *http.Request)  {
    log.Println("Got request!")
    client := &http.Client{}
    req, _ := http.NewRequest("POST", "https://andruxnet-random-famous-quotes.p.mashape.com/cat=movies", nil)
    req.Header.Set("X-Mashape-Key", "e3PM6jUqiYmsh3RbYScbSI8NgkpUp1hLBBXjsnJ6kJq7LUOnU3")
    resp, err := client.Do(req);
    //resp, err := http.Get("https://google.com")
    if err != nil {
        fmt.Fprintf(responseWrite, "Error connecting to InspireMe server %s", err)
    }  else {
        
        
        htmlData, _ := ioutil.ReadAll(resp.Body)
        
        fmt.Fprintf(responseWrite, string(htmlData))   
    }
}