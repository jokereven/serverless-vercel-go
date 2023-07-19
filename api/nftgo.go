package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func GetEthUsdRate() (string, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.nftgo.io/api/v1/asset/eth-usd-rate", nil)
    if err != nil {
        return "", err
    }
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}
