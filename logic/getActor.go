package logic

import (
    "github.com/bitly/go-simplejson"
)

func GetActor(ch chan <- *simplejson.Json) {
    url := "Movie/GetActor"
    data := make(map[string]string)
    data["PageIndex"] = "1"
	data["PageSize"] = "50"

    SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}