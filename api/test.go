package main

import (
    "fmt"
    "github.com/mozillazg/request"
    "net/http"
)

func main() {
    c := new(http.Client)
    req := request.NewRequest(c)
    req.Data = map[string]string{
        "PageIndex": "1",
        "PageSize":   "50",
    }
    resp, err := req.Post(Host + "Movie/GetClass")
    j, err := resp.Json()

    if err != nil {
        
    }

    fmt.Println(j)
}