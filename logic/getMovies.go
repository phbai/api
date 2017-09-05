package logic

import (
    "github.com/bitly/go-simplejson"
)

func GetMovies(ch chan <- *simplejson.Json) {
    url := "Movie/GetMovies"
    data := make(map[string]string)
    data["PageIndex"] = "1"
    data["PageSize"] = "1"
    data["Type"] = "1"
    data["ID"] = "-1"
    data["Data"] = ""

    SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}