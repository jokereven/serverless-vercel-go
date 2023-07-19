package nftgo

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("https://api.nftgo.io/api/v1/asset/eth-usd-rate")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var response Response
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
