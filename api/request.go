package main

import (
    "fmt"
    "github.com/mozillazg/request"
    "net/http"
    "github.com/bitly/go-simplejson"
)

func SendRequest(url string, data map[string]string, cb func(j *simplejson.Json)) {
    c := new(http.Client)
    req := request.NewRequest(c)
    // req.Headers = map[string]string{
    //     "Access-token": token,
    // }

    req.Data = data
    resp, err := req.Post(Host + url)
    j, err := resp.Json()

    if err != nil {
        fmt.Println(err)
    }

    cb(j);
    defer resp.Body.Close()  // Don't forget close the response body
}

func main() {
    // c := new(http.Client)
    // req := request.NewRequest(c)
    // req.Data = map[string]string{
    //     "PageIndex": "1",
    //     "PageSize":   "50",
    // }
    // resp, err := req.Post(Host + "Movie/GetClass")
    // j, err := resp.Json()

    // if err != nil {
        
    // }
    // var j *simplejson.Json = nil;
    GetChannel();
    
    // fmt.Println(j)
}