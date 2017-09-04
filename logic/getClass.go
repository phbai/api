package logic

import (
    "fmt"
    "github.com/bitly/go-simplejson"
)

func GetClass() {
    url := "Movie/GetClass"
    data := make(map[string]string)
    data["PageIndex"] = "1"
    data["PageSize"] = "50"

    SendRequest(url, data, func(j *simplejson.Json) {
        fmt.Println(j);
    })
}