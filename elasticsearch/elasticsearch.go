package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"encoding/json"
)

func panico(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	url := "http://localhost:9200/_aliases?pretty=1"
	
	res, err := http.Get(url)
	panico(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    panico(err)
    
    fmt.Printf("%s\n", body)
	// ...
}