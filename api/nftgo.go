package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.nftgo.io/api/v1/asset/eth-usd-rate", nil)
    if err != nil {
        panic(err)
    }
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(string(body))
}
