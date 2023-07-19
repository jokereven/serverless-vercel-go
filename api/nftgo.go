package nftgo

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    ErrorCode int     `json:"errorCode"`
    Data      Data    `json:"data"`
}

type Data struct {
    Rate float64 `json:"rate"`
}

func GetEthUsdRate(w http.ResponseWriter, r *http.Request) (float64, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.nftgo.io/api/v1/asset/eth-usd-rate", nil)
    if err != nil {
        return 0, err
    }
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }
    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        return 0, err
    }
    return response.Data.Rate, nil
}
