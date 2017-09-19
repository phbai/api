package util

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
